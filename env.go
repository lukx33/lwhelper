package lwhelper

import "syscall"

func GetEnv(key, def string) func() string {

	tmp, _ := syscall.Getenv(key)
	if tmp == "" {
		tmp = def
	}

	return func() string {
		return tmp
	}
}

func Mandatory(s string) string {
	if s == "" {
		panic("config is empty")
	}
	return s
}
