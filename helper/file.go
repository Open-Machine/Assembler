package helper

import (
	"path"
	"strings"
)

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, path.Ext(fileName))
}

func FileExtension(fileName string) string {
	return path.Ext(fileName)
}
