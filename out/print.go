package out

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(in interface{}) {

	fmt.Println(Trace(3)[0])

	// ---

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
