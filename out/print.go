package out

import (
	"encoding/json"
	"fmt"
)

var Debug = false
var DoPrint = true

func PrintJSON(in interface{}) {

	if !DoPrint {
		return
	}

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

	// fmt.Println(reflect.TypeOf(in))

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
