package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLibraryHandler(c *gin.Context) {

	libraryData, err := helper.GetLibraryFromYamlFile()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Library file not found",
		})
		return
	}

	c.JSON(http.StatusOK, libraryData)
}
