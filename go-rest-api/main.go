package main

import (
	"fmt"
	"net/http"

	"localflix/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	serveApplication()
}

func healthCheckHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"healthy": true,
	})
}

func serveApplication() {
	// create new gin router
	router := gin.Default()

	// Enable CORS
	router.Use(middleware.AllowAllCORS())

	router.GET("/", healthCheckHandler)

	router.Static("/assets", "./assets")

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
