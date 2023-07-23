package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLibraryHandler(c *gin.Context) {

	libraryData, err := helper.LoadLibrary()
	if err != nil {
		// Multiple possible causes for failure to loading the library (ie. corrupted file).
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error loading library",
		})
		return
	}

	c.JSON(http.StatusOK, libraryData)
}
