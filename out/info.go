package out

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/lukx33/lwhelper/result"
)

type StructS struct {
	Step []StepS           `json:"step" cbor:",omitempty"`
	Vars map[string]string `json:"vars" cbor:",omitempty"`
}

type StepS struct {
	Result result.CodeT `json:"result"`
	Trace  []string     `json:"trace"`
}

type Info interface {
	// CatchError(err error) bool
	// InfoMessage() string

	// Read:
	NotValid() bool
	InfoLastResult() result.CodeT
	// InfoTraces() []StepS
	// InfoLastTrace() StepS
	InfoJSON() []byte
	// InfoPrint()
	InfoFatal()

	// Write:
	InfoAddStep(result result.CodeT, skipFrames int) Info
	// InfoAddCause(parent Info) Info
	InfoAddVar(name string, value any) Info
}

// ---

func Trace(skip int) []string {
	trace := []string{}
	pc := make([]uintptr, 40)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, isMore := frames.Next()

		prefix := " "
		for k := range framesLessImportant {
			if strings.Contains(frame.File, k) {
				prefix = "*"
				break
			}
		}

		// trace = append(trace, fmt.Sprintf("%s %s:%d", prefix, frame.File, frame.Line))
		if prefix == " " {
			trace = append(trace, fmt.Sprintf("%s:%d", frame.File, frame.Line))
		}

		if !isMore {
			break
		}
	}
	return trace
}

var framesLessImportant = map[string]bool{
	"github.com/lukx33/lwhelper": true,
	"/timoni/timoni/lwhelper/":   true,
	"/golang/src/runtime/":       true,
	"/golang/src/net/http/":      true,
	"/global/result.go":          true,
	"/fyne/widget/":              true,
	"/fyne/internal/":            true,
}

// func ErrorHandler[T any](out T, err error) (*ResultS, T) {
// 	if err != nil {
// 		return Error(err, out), out
// 	}
// 	return Success(), out
// }

// ---
