package lwhelper

import (
	"fmt"
	"strings"
)

func QueryIn(key string, list []string) string {
	return fmt.Sprintf("%s IN (\"%s\")", key, strings.Join(list, "\",\""))
}
