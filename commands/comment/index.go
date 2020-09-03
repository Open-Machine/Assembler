package comment

import (
	"strings"
)

func RemoveComment(line string) string {
	index := strings.Index(line, "#")
	if index < 0 {
		return line
	}
	return line[:index]
}
