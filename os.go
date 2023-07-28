package lwhelper

import "os"

func ReadFile(filepath string) string {
	return string(Must(os.ReadFile(filepath)))
}
