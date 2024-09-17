package out

import (
	"encoding/json"
	"fmt"
	"os"
)

var Debug = os.Getenv("LW_OUT_DEBUG") == "true"

func PrintJSON(in interface{}) {

	if Debug {
		fmt.Println(Trace(3)[0])
	}

	// ---

	if !Debug {

		switch v := in.(type) {
		case *StructS:
			msg := v.Vars["msg"]
			if v.Vars["err"] != "" {
				msg = "ERROR: " + v.Vars["err"]
			}
			if msg != "" {
				fmt.Println(msg)
			}
			return
		}
	}

	buf, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(buf))
}

// func Must[T any](out T, err error) T {
// 	if err != nil {
// 		panic(err)
// 	}
// 	return out
// }
