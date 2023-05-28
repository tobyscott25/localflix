package handlers

import (
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetVideoListHandler(c *gin.Context) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")
	files := helper.GetFiles(libraryLocation)
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
