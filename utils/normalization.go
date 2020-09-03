package utils

import (
	"regexp"
	"strings"
)

func LineNormalization(line string) string {
	withoutNewLine := removeNewLine(line)
	withoutComment := removeComment(withoutNewLine)
	return removeUnecessarySpaces(withoutComment)
}

func removeNewLine(line string) string {
	lineWithoutEndingUnix := strings.TrimSuffix(line, "\n")
	lineWithoutEndingUnixAndWindows := strings.TrimSuffix(lineWithoutEndingUnix, "\r")
	return lineWithoutEndingUnixAndWindows
}

func removeUnecessarySpaces(line string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(line, " ")
	return strings.TrimSpace(str)
}

func removeComment(line string) string {
	index := strings.Index(line, "#")
	if index < 0 {
		return line
	}
	return line[:index]
}
