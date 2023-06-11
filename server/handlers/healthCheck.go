package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {

	if os.Getenv("LIBRARY_LOCATION") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"healthy": false,
			"message": "LIBRARY_LOCATION is not set",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"healthy":         true,
		"libraryLocation": os.Getenv("LIBRARY_LOCATION"),
	})
}
