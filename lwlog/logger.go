package lwlog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

//

func NewLogger() *slog.Logger {
	return slog.New(&logBufferHandler{
		handler: slog.NewJSONHandler(os.Stdout, nil),
	}).With("instance", os.Getenv("HOSTNAME"))
}

//

var (
	loggerWebSocketConnected = false
	loggerBuffer             = make(chan []byte, 2)
)

type logBufferHandler struct {
	handler slog.Handler
}

//

func (h *logBufferHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// return h.handler.Enabled(ctx, level)
	return true
}

func (h *logBufferHandler) Handle(ctx context.Context, r slog.Record) error {

	if loggerWebSocketConnected {

		buf, _ := json.Marshal(r)
		loggerBuffer <- buf
		// fmt.Println("delay:", time.Since(r.Time))
	}

	// Przekaż do oryginalnego handlera
	return h.handler.Handle(ctx, r)
}

func (h *logBufferHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logBufferHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *logBufferHandler) WithGroup(name string) slog.Handler {
	return &logBufferHandler{handler: h.handler.WithGroup(name)}
}

//

// HttpLogsHandler obsługuje WebSocket dla logów
func HttpLogsHandler(w http.ResponseWriter, r *http.Request) {

	if loggerWebSocketConnected {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Status": "error", "Message": "logsHandler already in use"}`)
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Ustawienia keepalive
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection", "error", err)
		return
	}
	fmt.Println("WebSocket connection established for logs")
	loggerWebSocketConnected = true

	// Ustawienia keepalive dla połączenia
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		return nil
	})

	// Goroutine do wysyłania pingów
	pingTicker := time.NewTicker(3 * time.Second)

	if err := conn.WriteMessage(websocket.TextMessage, []byte("WebSocket connection established for logs")); err != nil {
		fmt.Println("Failed to send logs", "error", err)
		return
	}

	defer func() {
		pingTicker.Stop()
		conn.Close()
		loggerWebSocketConnected = false
		fmt.Println("WebSocket disconnection for logs")
	}()

	// Goroutine do obsługi wiadomości od klienta
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("WebSocket error: %v\n", err)
				}
				return
			}
		}
	}()

	for {
		select {
		case msgBytes := <-loggerBuffer:
			if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
				fmt.Println("Failed to send logs", "error", err)
				return
			}
		case <-pingTicker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println("Failed to send ping", "error", err)
				return
			}
		case <-done:
			return
		}
	}
}
