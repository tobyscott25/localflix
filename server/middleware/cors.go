package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowAllCORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Allow all origins
		cors.Default()(c)

		// Handle OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue processing the request
		c.Next()
	}
}
