package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {

	if os.Getenv("LF_LIBRARY_LOCATION") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"healthy":    false,
			"testingair": "it works!",
			"message":    "LF_LIBRARY_LOCATION is not set",
		})
		return
	}

	// Add check for SQLite database file

	c.JSON(http.StatusOK, gin.H{
		"healthy":         true,
		"libraryLocation": os.Getenv("LF_LIBRARY_LOCATION"),
	})
}
