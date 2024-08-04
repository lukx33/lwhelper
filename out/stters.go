package out

import (
	"fmt"

	"github.com/lukx33/lwhelper/result"
)

// ---
// setters

func (info *StructS) InfoAddStep(result result.CodeT, skipFrames int) Info {

	info.Step = append(info.Step, StepS{
		Result: result,
		Trace:  Trace(skipFrames),
	})
	return info
}

// func (info *StructS) InfoAddCause(parent Info) Info {

// 	info.Step = append(info.Step, parent.InfoTraces()...)
// 	info.Step = append(info.Step, StepS{
// 		Result:  parent.InfoResult(),
// 		Message: "AddCause",
// 		Trace:   Trace(0),
// 	})
// 	info.Result = parent.InfoResult()
// 	info.InfoPrint()
// 	return info
// }

func (info *StructS) InfoAddVar(name string, value any) Info {

	if info.Vars == nil {
		info.Vars = map[string]string{}
	}
	info.Vars[name] = fmt.Sprint(value)

	return info
}
