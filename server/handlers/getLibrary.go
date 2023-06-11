package handlers

import (
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLibraryHandler(c *gin.Context) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	// Currently reading the actual videos, need to change this to read from the yaml file.
	videosList := helper.GetAllVideosInDirectory(libraryLocation)

	c.JSON(http.StatusOK, gin.H{
		"files": videosList,
	})
}
