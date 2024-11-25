package lwhelper

import (
	"fmt"
	"math/rand/v2"
	"runtime/debug"
	"strings"
	"time"

	"github.com/barkimedes/go-deepcopy"
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

func SleepRandomSec(min, max int) {
	if min <= 0 {
		min = 2
	}
	if min > max {
		max = min + 5
	}

	time.Sleep(time.Duration(rand.IntN(max-min)+min) * time.Second)
}

func DeepCopy[T any](src T) *T {
	copy, err := deepcopy.Anything(&src)
	if err != nil {
		return nil
	}
	return copy.(*T)
}
