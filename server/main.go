package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"localflix/server/helper"
	"localflix/server/middleware"

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

type FileInfoData struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Path string `json:"path"`
	// Mode    os.FileMode `json:"mode"`
	ModTime string `json:"lastModified"`
	// IsDir   bool        `json:"isDir"`
}

func getFileListHandler(c *gin.Context) {

	files := getFiles("assets")

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func getFiles(dirPath string) []FileInfoData {
	var files []FileInfoData

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		fileData := FileInfoData{
			Name:    info.Name(),
			Size:    helper.ByteCountSI(info.Size()),
			Path:    "/" + path,
			ModTime: info.ModTime().Format(time.RFC3339),
			// Mode:    info.Mode(),
			// IsDir:   info.IsDir(),
		}

		files = append(files, fileData)
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
