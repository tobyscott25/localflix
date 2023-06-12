package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideoDetailsHandler(c *gin.Context) {

	id := c.Param("id")

	library, err := helper.LoadLibraryFromYamlFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Missing Library file",
		})
		return
	}

	videoDetails := helper.GetVideoDetailsByID(*library, id)

	if videoDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Video not found",
		})
		return
	}

	c.JSON(http.StatusOK, videoDetails)
}
