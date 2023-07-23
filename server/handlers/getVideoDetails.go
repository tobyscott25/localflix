package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideoDetailsHandler(c *gin.Context) {

	id := c.Param("id")
	videoDetails := helper.GetVideoDetailsByID(id)

	if videoDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Video not found",
		})
		return
	}

	c.JSON(http.StatusOK, videoDetails)
}
