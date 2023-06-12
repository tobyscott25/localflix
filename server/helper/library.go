package helper

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type LibraryData struct {
	Version string         `json:"version"`
	Videos  []FileInfoData `json:"videos"` // Currently only videos are supported
}

func GetLibraryFromYamlFile() (*LibraryData, error) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	libraryFile, err := os.Open(libraryLocation + "/localflix-library.yaml")
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{
		// 	"error": "Library file not found",
		// })
		return nil, err
	}
	defer libraryFile.Close()

	libraryYamlData, err := io.ReadAll(libraryFile)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	var library LibraryData
	err = yaml.Unmarshal(libraryYamlData, &library)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return &library, nil
}

func GetVideoDetailsByID(library LibraryData, id string) *FileInfoData {
	for _, fileInfo := range library.Videos {
		if fileInfo.ID == id {
			return &fileInfo
		}
	}
	return nil
}
