package helper

import (
	"fmt"
	"localflix/server/model"
	"localflix/server/validation"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetAllVideosInDirectory(dirPath string) []model.Video {
	var videosArray []model.Video

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		if !validation.IsVideoFile(info.Name()) {
			return nil // Skip non-video files
		}

		videoInfo := model.Video{
			Title:          RemoveFilenameExtension(info.Name()), // use the filename (without extension) as the default title
			FileName:       info.Name(),
			FileSize:       HumanReadableFileSize(info.Size()),
			LastModified:   info.ModTime().Format(time.RFC3339),
			ChecksumSHA256: CalculateSHA256Checksum(path),
		}

		videosArray = append(videosArray, videoInfo)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}

	return videosArray
}

func RemoveFilenameExtension(filename string) string {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 {
		return filename
	}
	return filename[:lastDotIndex]
}
