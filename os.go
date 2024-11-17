package lwhelper

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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

func FileHasThisLine(filePath, line string) bool {

	buf, _ := os.ReadFile(filePath)
	return StringHasThisLine(string(buf), line)
}

func StringHasThisLine(s, line string) bool {

	line = strings.TrimSpace(line)

	for _, l := range strings.Split(s, "\n") {
		l = strings.TrimSpace(l)
		if l == line {
			return true
		}
	}

	return false
}

func FileFind(startPath, nameSuffix string, fullPath bool) ([]string, error) {

	if startPath == "" {
		return nil, errors.New("startPath is empty")
	}

	if startPath[len(startPath)-1] != '/' {
		startPath += "/"
	}

	res := []string{}

	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), nameSuffix) {

			if fullPath {
				res = append(res, path)

			} else {
				res = append(res, strings.TrimPrefix(path, startPath))
			}
		}
		return nil
	})

	return res, err
}
