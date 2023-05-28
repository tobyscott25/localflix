package helper

import (
	"fmt"
	"localflix/server/validation"
	"os"
	"path/filepath"
	"time"
)

type FileInfoData struct {
	Name        string `json:"name"`
	Size        string `json:"size"`
	Path        string `json:"path"`
	ModTime     string `json:"lastModified"`
	ChecksumSHA string `json:"checksum"`
}

func GetFiles(dirPath string) []FileInfoData {
	var files []FileInfoData

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

		fileData := FileInfoData{
			Name:        info.Name(),
			Size:        HumanReadableFileSize(info.Size()),
			Path:        "/" + path,
			ModTime:     info.ModTime().Format(time.RFC3339),
			ChecksumSHA: CalculateSHA256Checksum(path),
		}

		files = append(files, fileData)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}

	return files
}

func LookupFileByChecksum(checksum string) (string, error) {

	// Walk through the files to find a match with the given checksum
	// Optimise this by adding a library file that stores all checksums

	var filePath string
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	err := filepath.Walk(libraryLocation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileChecksum := CalculateSHA256Checksum(path)
			if fileChecksum == checksum {
				filePath = path
				return filepath.SkipDir
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if filePath == "" {
		return "", os.ErrNotExist
	}

	return filePath, nil
}
