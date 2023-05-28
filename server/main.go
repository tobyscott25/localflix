package main

import (
	"fmt"
	"os"

	"localflix/server/handlers"
	"localflix/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")
	serveApplication(libraryLocation)
}

func serveApplication(libraryLocation string) {

	fmt.Println("Library Path:", libraryLocation)

	router := gin.Default()
	router.Use(middleware.AllowAllCORS())

	router.GET("/", handlers.HealthCheckHandler)
	router.GET("/files", handlers.GetVideoListHandler)
	router.GET("/files/checksum/:checksum", handlers.GetVideoDetailsHandler)
	router.Static("/assets", libraryLocation)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Failed to start the server: %v", err)
		return
	}

	fmt.Println("Server running on port 8080")
}
