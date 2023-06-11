package helper

import (
	"fmt"
	"localflix/server/validation"
	"os"
	"path/filepath"
	"time"
)

type FileInfoData struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Size         string `json:"size"`
	Path         string `json:"path"`
	LastModified string `json:"lastModified"`
	ChecksumSHA  string `json:"checksum"`
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
			ID:           checksum, // temporarily use checksum as ID - this won't scale well.
			Name:         info.Name(),
			Size:         HumanReadableFileSize(info.Size()),
			Path:         path,
			LastModified: info.ModTime().Format(time.RFC3339),
			ChecksumSHA:  checksum,
		}

		videosArray = append(videosArray, videoInfo)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}

	return videosArray
}
