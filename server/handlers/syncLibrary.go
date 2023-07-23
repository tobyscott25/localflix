package handlers

import (
	"fmt"
	"localflix/server/database"
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SyncLibraryHandler(c *gin.Context) {

	version := os.Getenv("localflixSemanticVersion")
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	videosList := helper.GetAllVideosInDirectory(libraryLocation)

	dbVideosList := helper.GetAllDbVideosInDirectory(libraryLocation)

	for _, video := range dbVideosList {
		fmt.Println("Syncing video to DB:", video.Title)
		result := database.Database.Create(&video)
		fmt.Println(result.Error)
	}

	newLibrary := helper.LibraryData{
		Version: version,
		Videos:  videosList,
	}

	err := helper.WriteLibraryToYamlFile(newLibrary)
	if err != nil {
		fmt.Println("Failed to sync library file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sync library file",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
