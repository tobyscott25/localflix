package helper

import (
	"os"
	"path/filepath"
)

// func LookupFileByName(checksum string) (string, error) {

// 	// Walk through the files to find a match with the given checksum
// 	// Optimise this by adding a library file that stores all checksums

// 	var filePath string
// 	libraryLocation := os.Getenv("LIBRARY_LOCATION")

// 	err := filepath.Walk(libraryLocation, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if !info.IsDir() {
// 			fileChecksum := CalculateSHA256Checksum(path)
// 			if fileChecksum == checksum {
// 				filePath = path
// 				return filepath.SkipDir
// 			}
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	if filePath == "" {
// 		return "", os.ErrNotExist
// 	}

// 	return filePath, nil
// }

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
