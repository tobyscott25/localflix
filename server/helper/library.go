package helper

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type LibraryData struct {
	Version string         `json:"version"`
	Videos  []FileInfoData `json:"videos"` // Currently only videos are supported
}

func LoadLibraryFromYamlFile() (*LibraryData, error) {

	libraryFileLocation := os.Getenv("LIBRARY_LOCATION") + "/localflix-library.yaml"

	libraryFile, err := os.Open(libraryFileLocation)
	if err != nil {
		fmt.Printf("Failed to open library file: %v", err)
		return nil, err
	}
	defer libraryFile.Close()

	libraryYamlData, err := io.ReadAll(libraryFile)
	if err != nil {
		fmt.Printf("Failed to read library file: %v", err)
		return nil, err
	}

	var library LibraryData
	err = yaml.Unmarshal(libraryYamlData, &library)
	if err != nil {
		fmt.Printf("Failed to unmarshal library file: %v", err)
		return nil, err
	}

	return &library, nil
}

func SyncLibraryFile(library LibraryData) error {

	libraryFileLocation := os.Getenv("LIBRARY_LOCATION") + "/localflix-library.yaml"

	err := WriteYAMLFile(libraryFileLocation, library)
	if err != nil {
		fmt.Printf("Failed to write YAML file: %v", err)
		return err
	}

	return nil
}

func GetVideoDetailsByID(library LibraryData, id string) *FileInfoData {
	for _, fileInfo := range library.Videos {
		if fileInfo.ID == id {
			return &fileInfo
		}
	}

	return nil // Return nil if no videos match the ID
}
