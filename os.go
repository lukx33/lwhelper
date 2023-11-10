package lwhelper

import "os"

func ReadFile(filepath string) string {
	return string(Must(os.ReadFile(filepath)))
}

func OsArg(idx int) string {
	if len(os.Args) <= idx {
		return append(os.Args, make([]string, idx)...)[idx]
	}
	return os.Args[idx]
}
