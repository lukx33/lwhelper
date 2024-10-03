package lwhelper

import (
	"cmp"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

func IDshort() string {
	return strings.ReplaceAll(
		uuid.New().String()[:10],
		"-", "",
	)
}

func ID() string {
	return strings.ReplaceAll(
		uuid.New().String(),
		"-", "",
	)
}

func IDlong() string {
	return strings.ReplaceAll(
		uuid.New().String()+uuid.New().String(),
		"-", "",
	)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N}]+`)
var plReplacer = strings.NewReplacer(
	"ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z",
	"Ą", "A", "Ć", "C", "Ę", "E", "Ł", "L", "Ń", "N", "Ó", "O", "Ś", "S", "Ż", "Z", "Ź", "Z",
)

func KeyString(s string) string {
	// pozostawia tylko male literki, cyfry i -
	key := strings.ToLower(s)
	key = plReplacer.Replace(key)
	key = nonAlphanumericRegex.ReplaceAllString(key, "-")
	key = strings.Trim(key, "-")
	return key
}

func KeyString2(s string) string {
	// pozostawia tylko literki i cyfry
	key := plReplacer.Replace(s)
	key = nonAlphanumericRegex.ReplaceAllString(key, "")
	return key
}

func CleanString(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) || r == '\n' {
			return r
		}
		return -1
	}, str)
}

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

func DateStringToTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}

func StringToTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}

func UnixToString(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04")
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
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
