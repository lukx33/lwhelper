package out

import (
	"fmt"
	"runtime"
	"strings"
)

type DontUseMeInfoS struct {
	Trace  []traceS          `json:"trace"`
	Vars   map[string]string `json:"vars"`
	Result ResultT           `json:"result"`
}

type traceS struct {
	Result    ResultT  `json:"result"`
	Message   string   `json:"message"`
	Traceback []string `json:"traceback"`
}

type Info interface {
	InfoSetSuccess() Info
	InfoSetError(err error) Info
	InfoSetResult(result ResultT) Info
	InfoAddVar(name string, value any) Info

	NotValid() bool
	InfoMessage() string
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
		Trace: []traceS{{
			Result:    Success,
			Traceback: Trace(0),
		}},
		Result: Success,
	}
}

func NewSuccessMsg(msg string) Info {

	r := &DontUseMeInfoS{
		Trace: []traceS{{
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
		Trace: []traceS{{
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
		Trace: []traceS{{
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
		Trace: []traceS{{
			Result:    NotFound,
			Traceback: Trace(0),
		}},
		Result: NotFound,
	}

	r.InfoPrint()
	return r
}

func NewForbidden() Info {

	r := &DontUseMeInfoS{
		Trace: []traceS{{
			Result:    Forbidden,
			Traceback: Trace(0),
		}},
		Result: Forbidden,
	}

	r.InfoPrint()
	return r
}

// ---
// setters

func (r *DontUseMeInfoS) InfoSetSuccess() Info {

	r.Trace = append(r.Trace, traceS{
		Result:    Success,
		Traceback: Trace(0),
	})
	r.Result = Success

	return r
}

func (r *DontUseMeInfoS) InfoSetError(err error) Info {

	if err == nil {
		return r
	}

	r.Trace = append(r.Trace, traceS{
		Result:    InternalServerError,
		Message:   err.Error(),
		Traceback: Trace(0),
	})
	r.Result = InternalServerError

	r.InfoPrint()
	return r
}

func (r *DontUseMeInfoS) InfoSetResult(result ResultT) Info {

	r.Trace = append(r.Trace, traceS{
		Result:    result,
		Traceback: Trace(0),
	})
	r.Result = result

	r.InfoPrint()
	return r
}

func (r *DontUseMeInfoS) InfoAddVar(name string, value any) Info {

	if r.Vars == nil {
		r.Vars = map[string]string{}
	}
	r.Vars[name] = fmt.Sprint(value)

	return r
}

// ---
// checkers

func (r *DontUseMeInfoS) NotValid() bool {

	if r == nil {
		return true
	}

	if r.Result == Success {
		return false
	}

	return true
}

func (r *DontUseMeInfoS) InfoMessage() string {
	return fmt.Sprintf("Result: %d %s %s", r.Result, r.Vars["error"], r.Vars["msg"])
}

func (r *DontUseMeInfoS) InfoPrint() {
	// traceStr := strings.Join(r.Traceback, "\n\t")

	// fmt.Printf("%s\n\t%s\n",
	// 	r.String(),
	// 	traceStr,
	// )
	PrintJSON(r)
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
	"/lwhelper/out/":       true,
	"/golang/src/runtime/": true,
	"/global/result.go":    true,
	"/fyne/widget/":        true,
	"/fyne/internal/":      true,
}

// func ErrorHandler[T any](out T, err error) (*ResultS, T) {
// 	if err != nil {
// 		return Error(err, out), out
// 	}
// 	return Success(), out
// }
