package handlers

import (
	"fmt"
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideoDetailsHandler(c *gin.Context) {

	id := c.Param("id")

	library, err := helper.GetLibraryFromYamlFile()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Library file not found",
		})
		return
	}

	fmt.Println("ID:", id)
	videoDetails := helper.GetVideoDetailsByID(*library, id)

	fmt.Println("Video Details:", videoDetails)

	if videoDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Video not found",
		})
		return
	}

	c.JSON(http.StatusOK, videoDetails)
}
