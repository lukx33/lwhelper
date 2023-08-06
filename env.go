package lwhelper

import "syscall"

func GetEnv(key, def string, mandatory bool) func() string {

	tmp, _ := syscall.Getenv(key)
	if tmp == "" {
		if def == "" && mandatory {
			panic("config `" + key + "` is empty")
		}
		tmp = def
	}

	return func() string {
		return tmp
	}
}
