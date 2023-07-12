package main

import (
	"fmt"
	"os"

	"localflix/server/handlers"
	"localflix/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("localflixSemanticVersion", "0.1.0")

	libraryLocation := os.Getenv("LIBRARY_LOCATION")
	serveApplication(libraryLocation)
}

func serveApplication(libraryLocation string) {

	fmt.Println("Library Path:", libraryLocation)

	router := gin.Default()
	router.Use(middleware.AllowAllCORS())

	router.GET("/", handlers.HealthCheckHandler)
	router.GET("/library", handlers.GetLibraryHandler)
	router.POST("/library/sync", handlers.SyncLibraryHandler)
	router.GET("/library/videos/:id", handlers.GetVideoDetailsHandler)
	router.PUT("/library/videos/:id", handlers.UpdateVideoDetailsHandler)
	router.Static("/assets", libraryLocation)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Failed to start the server: %v", err)
		return
	}

	fmt.Println("Server running on port 8080")
}
