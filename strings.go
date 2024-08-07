package lwhelper

import (
	"regexp"
	"strconv"
	"strings"
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
}
