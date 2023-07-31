package out

import (
	"fmt"
	"runtime"
	"strings"
)

type DontUseMeInfoS struct {
	Trace  []TraceS          `json:"trace"`
	Vars   map[string]string `json:"vars"`
	Result ResultT           `json:"result"`
}

type TraceS struct {
	Result    ResultT  `json:"result"`
	Message   string   `json:"message"`
	Traceback []string `json:"traceback"`
}

type Info interface {
	CatchError(err error) bool
	NotValid() bool
	InfoAddTrace(result ResultT, msg string, skipFrames int)
	InfoAddCause(parent Info) Info
	InfoAddVar(name string, value any) Info
	InfoResult() ResultT
	InfoMessage() string
	InfoTrace() []TraceS
	InfoPrint()
}

type ResultT uint

const (
	// api / request / function:
	Unknown             ResultT = 0
	Success             ResultT = 1
	InternalServerError ResultT = 2
	ValidationError     ResultT = 3 // wrong input data, bad request
	Timeout             ResultT = 4
	Forbidden           ResultT = 5 // permission deny

	// resource / object:
	NotFound      ResultT = 20
	CreateError   ResultT = 21
	DeleteError   ResultT = 22
	AlreadyExists ResultT = 23

	// login / session:
	// InvalidEmail        ResultCode = 2
	// MissingTotpToken    ResultCode = 3
	// InvalidTotpToken    ResultCode = 4
	// InvalidLicenseKey   ResultCode = 5
	// SessionIDInvalid  ResultCode = 32
	// SessionIDNotFound ResultCode = 33

	FistLogin ResultT = 30
	// UserIsBlocked ResultCode = 31

	// InstalationNotFound ResultCode = 40
	// InstalationBlocked  ResultCode = 41

	// ActionNotFound       ResultCode = 60
	// ClusterConfigIsEmpty ResultCode = 61
)

// ---
// constructors

func NewSuccess() Info {

	return &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    Success,
			Traceback: Trace(0),
		}},
		Result: Success,
	}
}

func NewSuccessMsg(msg string) Info {

	r := &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    Success,
			Message:   msg,
			Traceback: Trace(0),
		}},
		Result: Success,
	}
	r.InfoPrint()
	return r
}

func NewError(err error) Info {

	if err == nil {
		return NewSuccess()
	}

	r := &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    InternalServerError,
			Message:   err.Error(),
			Traceback: Trace(0),
		}},
		Result: InternalServerError,
	}

	r.InfoPrint()
	return r
}

func NewErrorMsg(msg string) Info {

	r := &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    InternalServerError,
			Message:   msg,
			Traceback: Trace(0),
		}},
		Result: InternalServerError,
	}

	r.InfoPrint()
	return r
}

func NewNotFound() Info {

	r := &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    NotFound,
			Traceback: Trace(0),
		}},
		Result: NotFound,
	}

	r.InfoPrint()
	return r
}

func NewForbidden() Info {

	info := &DontUseMeInfoS{
		Trace: []TraceS{{
			Result:    Forbidden,
			Traceback: Trace(0),
		}},
		Result: Forbidden,
	}

	info.InfoPrint()
	return info
}

// ---
// setters

func (info *DontUseMeInfoS) InfoAddTrace(result ResultT, msg string, skipFrames int) {

	info.Trace = append(info.Trace, TraceS{
		Result:    result,
		Message:   msg,
		Traceback: Trace(skipFrames),
	})
	info.Result = result
}

func (info *DontUseMeInfoS) InfoAddCause(parent Info) Info {

	info.Trace = append(info.Trace, parent.InfoTrace()...)
	info.Trace = append(info.Trace, TraceS{
		Result:    parent.InfoResult(),
		Message:   "AddCause",
		Traceback: Trace(0),
	})
	info.Result = parent.InfoResult()
	info.InfoPrint()
	return info
}

func (info *DontUseMeInfoS) InfoAddVar(name string, value any) Info {

	if info.Vars == nil {
		info.Vars = map[string]string{}
	}
	info.Vars[name] = fmt.Sprint(value)

	return info
}

func (info *DontUseMeInfoS) CatchError(err error) bool {

	if err == nil {
		return false
	}

	info.InfoAddTrace(InternalServerError, err.Error(), 0)
	info.InfoPrint()
	return true
}

// ---
// checkers

func (info *DontUseMeInfoS) NotValid() bool {

	if info == nil {
		return true
	}

	if info.Result == Success {
		return false
	}

	return true
}

func (info *DontUseMeInfoS) InfoResult() ResultT {
	return info.Result
}

func (info *DontUseMeInfoS) InfoMessage() string {
	return fmt.Sprintf("Result: %d", info.Result)
	// TODO: info.Trace[*].Msg
}

func (info *DontUseMeInfoS) InfoTrace() []TraceS {
	return info.Trace
}

func (info *DontUseMeInfoS) InfoPrint() {
	// traceStr := strings.Join(r.Traceback, "\n\t")

	// fmt.Printf("%s\n\t%s\n",
	// 	r.String(),
	// 	traceStr,
	// )
	PrintJSON(info)
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
	"/golang/src/runtime/":       true,
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

func SetSuccess[T Info](info T) T {
	info.InfoAddTrace(Success, "", 0)
	return info
}

func CatchError[T Info](info T, err error) T {

	if err == nil {
		return info
	}

	info.InfoAddTrace(InternalServerError, err.Error(), 0)
	info.InfoPrint()
	return info
}
