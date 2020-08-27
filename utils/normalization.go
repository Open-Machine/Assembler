package utils

import (
	"regexp"
	"strings"
)

func LineNormalization(line string) string {
	withoutNewLine := RemoveNewLine(line)
	return removeUnecessarySpaces(withoutNewLine)
}

func RemoveNewLine(line string) string {
	lineWithoutEndingUnix := strings.TrimSuffix(line, "\n")
	lineWithoutEndingUnixAndWindows := strings.TrimSuffix(lineWithoutEndingUnix, "\r")
	return lineWithoutEndingUnixAndWindows
}

func removeUnecessarySpaces(line string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(line, " ")
	return strings.TrimSpace(str)
}
