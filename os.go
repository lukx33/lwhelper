package lwhelper

import "os"

func ReadFile(filepath string) string {
	return string(Must(os.ReadFile(filepath)))
}

func FileExist(filepath string) bool {
	if fs, err := os.Stat(filepath); err == nil {
		return !fs.IsDir()
	}
	return false
}

func OsArg(idx int) string {
	if len(os.Args) <= idx {
		return append(os.Args, make([]string, idx)...)[idx]
	}
	return os.Args[idx]
}

func OsArgs(idx int) []string {
	idx++
	if len(os.Args) <= idx {
		return append(os.Args, make([]string, idx)...)[:idx]
	}
	return os.Args[:idx]
}
