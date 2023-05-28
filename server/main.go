package main

import (
	"fmt"

	"localflix/server/handlers"
	"localflix/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.Use(middleware.AllowAllCORS())

	router.GET("/", handlers.HealthCheckHandler)
	router.GET("/files", handlers.GetFileListHandler)
	router.GET("/files/:fileName", handlers.GetFileDetailsHandler)
	router.GET("/files/checksum/:checksum", handlers.GetFileByChecksumHandler)
	router.Static("/assets", "./assets")

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Failed to start the server: %v", err)
		return
	}

	fmt.Println("Server running on port 8080")
}
