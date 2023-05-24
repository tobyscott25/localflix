package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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

func getFileListHandler(c *gin.Context) {

	files := getFiles("assets")

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func getFiles(dirPath string) []string {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}

	return files
}

func serveApplication() {

	router := gin.Default()

	router.Use(middleware.AllowAllCORS())

	router.GET("/", healthCheckHandler)
	router.GET("/files", getFileListHandler)
	router.Static("/assets", "./assets")

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
