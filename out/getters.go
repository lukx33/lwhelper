package out

import (
	"encoding/json"
	"os"

	"github.com/lukx33/lwhelper/result"
)

// ---
// getters / checkers

// func (info *StructS) infoSteps() []StepS {
// 	return info.Step
// }

func (info *StructS) infoLastTrace() StepS {
	if len(info.Step) == 0 {
		return StepS{}
	}
	return info.Step[len(info.Step)-1]
}

// ----------------------------------------------------
// exposed

func (info *StructS) InfoLastResult() result.CodeT {
	return info.infoLastTrace().Result
}

func (info *StructS) NotValid() bool {

	if info == nil {
		return true
	}

	return info.InfoLastResult() != result.Success
}

func (info *StructS) InfoJSON() []byte {
	buf, _ := json.MarshalIndent(info, "", "  ")
	return buf
}

func (info *StructS) InfoFatal() {
	if info.InfoLastResult() != result.Success {
		info.InfoAddVar("fatal", "true")
		PrintJSON(info)
		os.Exit(1)
	}
}

// func (info *StructS) InfoPrint() {
// 	fmt.Println(Trace(3)[0])
// 	fmt.Println(string(info.InfoJSON()))
// }
