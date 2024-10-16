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

	libraryLocation := os.Getenv("LF_LIBRARY_LOCATION")

	// Delete all entries in the 'videos' table
	// queryResult := database.Database.Delete(&model.Video{}) // This isn't allowed by GORM
	queryResult := database.Database.Exec("DELETE FROM videos")
	if queryResult.Error != nil {
		fmt.Println("Error deleting entries:", queryResult.Error)
	} else {
		fmt.Println("Deleted rows:", queryResult.RowsAffected)
	}

	videosList := helper.GetAllVideosInDirectory(libraryLocation)

	for _, video := range videosList {
		fmt.Println("Syncing video to DB:", video.Title)
		result := database.Database.Create(&video)
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusNoContent, nil)
}
