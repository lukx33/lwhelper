package lwhelper

import "os"

func ReadFile(filePath string) string {
	return string(Must(os.ReadFile(filePath)))
}

func FileExist(filePath string) bool {
	if fs, err := os.Stat(filePath); err == nil {
		return !fs.IsDir()
	}
	return false
}

func DirExist(dirPath string) bool {
	if fs, err := os.Stat(dirPath); err == nil {
		return fs.IsDir()
	}
	return false
}

func PathExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
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
