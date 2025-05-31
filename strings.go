package lwhelper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image/color"
	"regexp"
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
var reCleanUp = regexp.MustCompile(`[^a-zA-Z0-9\.\_\-]+`)

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

func KeyString3(s string) string {
	// pozostawia tylko male literki, cyfry i -
	s = strings.TrimSpace(plReplacer.Replace(s))
	return strings.Trim(reCleanUp.ReplaceAllString(s, "-"), "-")
}

func CleanString(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) || r == '\n' {
			return r
		}
		return -1
	}, str)
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

func GetMD5Hash(buf []byte) string {
	hash := md5.Sum(buf)
	return hex.EncodeToString(hash[:])
}

func GetSHA256(in []byte) string {
	var buf bytes.Buffer
	for i, f := range in {
		if i > 0 {
			fmt.Fprintf(&buf, ":")
		}
		fmt.Fprintf(&buf, "%02X", f)
	}
	return buf.String()
}

func StringSplit(s, delimiter string) []string {

	s = strings.TrimSpace(s)
	res := []string{}
	for _, x := range strings.Split(s, delimiter) {
		x = strings.TrimSpace(x)
		if x == "" {
			continue
		}
		res = append(res, x)
	}
	return res
}

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func HexColor(s string, opacity float64) string {
	c := color.RGBA{}

	switch len(s) {
	case 7:
		fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	}
	return fmt.Sprintf("rgba(%d, %d, %d, %.1f)", c.R, c.G, c.B, opacity)
}

//

func CleanInvisibleChars(s string) string {
	var builder strings.Builder

	for _, r := range s {
		if unicode.IsControl(r) && r != '\n' && r != '\t' {
			continue
		}
		if unicode.In(r, unicode.Cf, unicode.Zl, unicode.Zp) {
			continue
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

//

func StringSplitTwo(s, sep string) (string, string) {

	x := strings.SplitN(s, sep, 2)
	if len(x) < 2 {
		x = append(x, "")
	}
	return x[0], x[1]
}
