package lwhelper

import (
	"strings"

	"github.com/google/uuid"
)

func ID() string {
	return strings.ReplaceAll(
		uuid.New().String()+uuid.New().String(),
		"-", "",
	)
}
