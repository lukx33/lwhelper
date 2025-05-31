package lwhelper

import (
	"cmp"
	"math"
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return i

	// 	var n int64
	// 	fmt.Sscan(strings.TrimSpace(s), &n)
	// 	return n
}

func ToFloat64(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}

func Float64Round(num float64, precision int) float64 {

	round := func(num float64) int {
		return int(num + math.Copysign(0.5, num))
	}

	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func DivideToCeil[T int | int64](a, b T) T {

	if b == 0 {
		return 0
	}

	// math.Ceil(float64(a) / float64(b))
	return (a + b - 1) / b
}

func Sum[T cmp.Ordered](vec []T) T {
	var sum T
	for _, elt := range vec {
		sum = sum + elt
	}
	return sum
}

func Max[T cmp.Ordered](s []T) T {
	var res T
	for _, v := range s {
		if v > res {
			res = v
		}
	}
	return res
}

func Min(s []float64) float64 {
	res := math.MaxFloat64
	for _, v := range s {
		if v < res {
			res = v
		}
	}
	return res
}
