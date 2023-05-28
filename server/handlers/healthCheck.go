package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"healthy":         true,
		"libraryLocation": os.Getenv("LIBRARY_LOCATION"),
	})
}
