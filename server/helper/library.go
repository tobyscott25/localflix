package helper

import (
	"fmt"
	"localflix/server/database"
	"localflix/server/model"
	"os"
)

type LibraryData struct {
	Version string        `json:"version"`
	Videos  []model.Video `json:"videos"` // Currently only videos are supported
}

func LoadLibrary() (*LibraryData, error) {

	version := os.Getenv("LF_VERSION")
	videos := GetAllVideosInDB()

	return &LibraryData{
		Version: version,
		Videos:  videos,
	}, nil
}

func GetAllVideosInDB() []model.Video {
	var videos []model.Video
	database.Database.Find(&videos)
	return videos
}

func GetVideoDetailsByID(id string) *model.Video {

	var video model.Video
	err := database.Database.Where("id = ?", id).First(&video).Error

	if err != nil {
		fmt.Println("Error getting video details:", err)
		return nil // Return nil if no videos match the ID
	}

	fmt.Println("Video details (TITLE):", video.Title)

	return &video
}
