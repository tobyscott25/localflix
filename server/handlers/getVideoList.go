package handlers

import (
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetVideoListHandler(c *gin.Context) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")
	videosList := helper.GetAllVideos(libraryLocation)
	c.JSON(http.StatusOK, gin.H{
		"files": videosList,
	})
}
