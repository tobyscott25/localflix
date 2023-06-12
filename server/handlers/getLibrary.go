package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLibraryHandler(c *gin.Context) {

	libraryData, err := helper.LoadLibraryFromYamlFile()

	if err != nil {

		// There may be other possible causes for failure to loading the library (ie. corrupted file).
		// If error isn't nil, we should read the error to see what actually went wrong.
		// I'll add that error handling later. Just telling the user it doesn't exist for now.

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Library file not found",
		})
		return
	}

	c.JSON(http.StatusOK, libraryData)
}
