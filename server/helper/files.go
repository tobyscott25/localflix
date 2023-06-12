package helper

import (
	"fmt"
	"localflix/server/validation"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileInfoData struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Size           string `json:"size"`
	Path           string `json:"path"`
	LastModified   string `json:"last_modified"`
	ChecksumSHA256 string `json:"checksum_sha256"`
}

func GetAllVideosInDirectory(dirPath string) []FileInfoData {
	var videosArray []FileInfoData

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

		checksum := CalculateSHA256Checksum(path)

		videoInfo := FileInfoData{
			ID:             checksum,                             // temporarily use checksum as ID - this won't scale well.
			Title:          RemoveFilenameExtension(info.Name()), // use the filename (without extension) as the default title
			Size:           HumanReadableFileSize(info.Size()),
			Path:           "/assets/" + info.Name(),
			LastModified:   info.ModTime().Format(time.RFC3339),
			ChecksumSHA256: checksum,
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
