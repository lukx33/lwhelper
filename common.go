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

func Ptr[T any](v T) *T {
	return &v
}

var Colors = []string{
	"#36A2EB", "#FF6384", "#FFCE56", "#4BC0C0", "#9966FF", "#FF9F40",
	"#C9CBCF", "#2E2E2E", "#FF5733", "#33FF57", "#33FFF6", "#3375FF",
	"#8A33FF", "#FF33C7", "#FFC733", "#B833FF", "#33FF8A", "#FF3333",
	"#33FFFF", "#3333FF", "#1F77B4", "#FF7F0E", "#2CA02C", "#D62728",
	"#9467BD", "#8C564B", "#E377C2", "#7F7F7F", "#BCBD22", "#17BECF",
	"#FCE205", "#FF5722", "#009688", "#795548", "#607D8B", "#3F51B5",
	"#E91E63", "#9C27B0", "#673AB7", "#4CAF50", "#CDDC39", "#FFC107",
	"#00BCD4", "#8BC34A", "#FF9800", "#FFEB3B", "#9E9E9E", "#2196F3",
	"#03A9F4", "#C2185B", "#512DA8", "#304FFE", "#0288D1", "#388E3C",
	"#7B1FA2", "#E65100", "#BF360C", "#FF5252", "#448AFF", "#76FF03",
}
