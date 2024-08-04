package out

import (
	"github.com/lukx33/lwhelper/result"
)

// ---
// constructors

func Success() Info {

	return &StructS{
		Step: []StepS{{
			Result: result.Success,
			Trace:  Trace(0),
		}},
	}
}

func CheckError(err error) Info {

	if err == nil {
		return Success()
	}

	r := &StructS{
		Step: []StepS{{
			Result: result.Error,
			Trace:  Trace(0),
		}},
		Vars: map[string]string{
			"err": err.Error(),
		},
	}

	PrintJSON(r)
	return r
}

// func CheckErrorVars(err error, vars map[string]string) Info {

// 	if err == nil {
// 		return Success()
// 	}

// 	r := &StructS{
// 		Step: []StepS{{
// 			Result: result.Error,
// 			Trace:  Trace(0),
// 		}},
// 		Vars: vars,
// 	}
// 	if r.Vars == nil {
// 		r.Vars = map[string]string{}
// 	}
// 	r.Vars["err"] = err.Error()

// 	PrintJSON(r)
// 	return r
// }

func ErrorMsg(msg string) Info {

	r := &StructS{
		Step: []StepS{{
			Result: result.Error,
			Trace:  Trace(0),
		}},
		Vars: map[string]string{
			"err": msg,
		},
	}

	PrintJSON(r)
	return r
}

func DebugMsg(msg string) Info {

	r := &StructS{
		Step: []StepS{{
			Result: result.Debug,
			Trace:  Trace(0),
		}},
		Vars: map[string]string{
			"msg": msg,
		},
	}

	PrintJSON(r)
	return r
}

func Error(rc result.CodeT, vars map[string]string) Info {

	r := &StructS{
		Step: []StepS{{
			Result: rc,
			Trace:  Trace(0),
		}},
		Vars: vars,
	}

	PrintJSON(r)
	return r
}

// func Forbidden() Info {

// 	info := &StructS{
// 		Step: []StepS{{
// 			Result: result.Forbidden,
// 			Trace:  Trace(0),
// 		}},
// 	}

// 	PrintJSON(info)
// 	return info
// }

// func NeedRoot() Info {

// 	info := &StructS{
// 		Step: []StepS{{
// 			Result: result.NeedRoot,
// 			Trace:  Trace(0),
// 		}},
// 	}

// 	PrintJSON(info)
// 	return info
// }

// func Nil() Info {

// 	info := &StructS{
// 		Step: []StepS{{
// 			Result: result.Nil,
// 			Trace:  Trace(0),
// 		}},
// 	}

// 	PrintJSON(info)
// 	return info
// }

// --- for object:

func SuccessFor[T Info](info T) T {
	info.InfoAddStep(result.Success, 0)
	return info
}

func CheckErrorFor[T Info](info T, err error) T {

	if err == nil {
		if info.InfoLastResult() == result.Unknown {
			info.InfoAddStep(result.Success, 0)
		}
		return info
	}

	info.InfoAddStep(result.Error, 0)
	info.InfoAddVar("err", err.Error())
	PrintJSON(info)
	return info
}
