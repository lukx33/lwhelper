package lwhelper

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filePath string) string {

	buf, _ := os.ReadFile(filePath)
	return string(buf)
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

func FileFind(startPath, nameSuffix string, returnFullPath bool) ([]string, error) {

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

			if returnFullPath {
				res = append(res, path)

			} else {
				res = append(res, strings.TrimPrefix(path, startPath))
			}
		}
		return nil
	})

	return res, err
}

func DirSizeInMB(path string) (int64, error) {
	var sizeBytes int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			sizeBytes += info.Size()
		}
		return err
	})
	return int64(math.Ceil(float64(sizeBytes) / 1048576)), err
}
