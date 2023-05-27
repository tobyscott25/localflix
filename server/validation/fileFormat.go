package validation

import (
	"path/filepath"
	"strings"
)

func IsVideoFile(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	return ext == ".mp4" || ext == ".mov" || ext == ".avi" || ext == ".mkv"
}
