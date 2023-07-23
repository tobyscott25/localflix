package handlers

import (
	"fmt"
	"localflix/server/database"
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SyncLibraryHandlerv2(c *gin.Context) {

	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	// Delete all entries in the 'videos' table
	// result := database.Database.Delete(&model.Video{})
	result := database.Database.Exec("DELETE FROM videos")
	if result.Error != nil {
		// Handle error, if any
		fmt.Println("Error deleting entries:", result.Error)
	} else {
		// Print the number of rows deleted
		fmt.Println("Deleted rows:", result.RowsAffected)
	}

	dbVideosList := helper.GetAllDbVideosInDirectory(libraryLocation)

	for _, video := range dbVideosList {
		fmt.Println("Syncing video to DB:", video.Title)
		result := database.Database.Create(&video)
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusNoContent, nil)
}
