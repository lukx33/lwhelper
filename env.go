package lwhelper

import "syscall"

func GetEnv(key, def string) func() string {

	tmp, _ := syscall.Getenv(key)
	if tmp == "" {
		if def == "" {
			panic("config `" + key + "` is empty")
		}
		tmp = def
	}

	return func() string {
		return tmp
	}
}
