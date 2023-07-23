package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideoDetailsHandlerV2(c *gin.Context) {

	id := c.Param("id")
	videoDetails := helper.GetVideoDetailsByIDv2(id)

	if videoDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Video not found",
		})
		return
	}

	c.JSON(http.StatusOK, videoDetails)
}
