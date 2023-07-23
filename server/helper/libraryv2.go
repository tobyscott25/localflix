package helper

import (
	"fmt"
	"localflix/server/database"
	"localflix/server/model"
	"os"
)

type LibraryDataV2 struct {
	Version string        `json:"version"`
	Videos  []model.Video `json:"videos"` // Currently only videos are supported
}

func LoadLibraryV2() (*LibraryDataV2, error) {

	// libraryFileLocation := os.Getenv("LIBRARY_LOCATION") + "/localflix-library.yaml"
	version := os.Getenv("localflixSemanticVersion")
	videos := GetAllVideosInDB()

	return &LibraryDataV2{
		Version: version,
		Videos:  videos,
	}, nil
}

func GetAllVideosInDB() []model.Video {
	var videos []model.Video
	database.Database.Find(&videos)
	return videos
}

func GetVideoDetailsByIDv2(id string) *model.Video {

	var video model.Video
	err := database.Database.Where("id = ?", id).First(&video).Error

	if err != nil {
		fmt.Println("Error getting video details:", err)
		return nil // Return nil if no videos match the ID
	}

	fmt.Println("Video details (TITLE):", video.Title)

	return &video
}

func UpdateVideoDetailsByIDv2(library *LibraryData, id string, updatedFileInfo model.Video) error {
	for i, fileInfo := range library.Videos {
		if fileInfo.ID == id {

			// Update the specified fields of the model.Video object
			if updatedFileInfo.Title != "" {
				library.Videos[i].Title = updatedFileInfo.Title
			}
			if updatedFileInfo.Description != "" {
				library.Videos[i].Description = updatedFileInfo.Description
			}
			if updatedFileInfo.FileSize != "" {
				return fmt.Errorf("size cannot be updated")
			}
			if updatedFileInfo.FileName != "" {
				return fmt.Errorf("path cannot be updated")
			}
			if updatedFileInfo.LastModified != "" {
				return fmt.Errorf("LastModified cannot be updated manually")
			}
			if updatedFileInfo.ChecksumSHA256 != "" {
				return fmt.Errorf("ChecksumSHA256 cannot be updated manually")
			}

			return nil
		}
	}
	return nil
}
