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
	"#FF6384", // Czerwony różowy
	"#36A2EB", // Niebieski
	"#FFCE56", // Żółty
	"#4BC0C0", // Turkusowy
	"#9966FF", // Fioletowy
	"#FF9F40", // Pomarańczowy
	"#C9CBCF", // Szary jasny
	"#2E2E2E", // Ciemny szary
	"#FF5733", // Ceglasty
	"#33FF57", // Zielony jasny
	"#33FFF6", // Turkusowy jasny
	"#3375FF", // Niebieski mocny
	"#8A33FF", // Fioletowy głęboki
	"#FF33C7", // Różowy intensywny
	"#FFC733", // Złoty
	"#B833FF", // Purpurowy
	"#33FF8A", // Miętowy
	"#FF3333", // Intensywny czerwony
	"#33FFFF", // Cyan jasny
	"#3333FF", // Granatowy
}
