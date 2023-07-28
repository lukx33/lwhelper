package lwhelper

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Must[T any](out T, err error) T {
	if err != nil {
		panic(err)
	}
	return out
}

func PanicHandler() {
	if err := recover(); err != nil {
		fmt.Println("Panic:", err)
		_, stack, _ := strings.Cut(string(debug.Stack()), "panic")
		fmt.Println("Stack:", stack)
	}
}
