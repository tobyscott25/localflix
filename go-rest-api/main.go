package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func getFileListHandler(c *gin.Context) {

	// get the list of files in the assets directory
	// and return them as a JSON array
	files, err := getFiles("assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func getFiles(dir string) ([]string, error) {

	files, err := ioutil.ReadDir("assets")
	if err != nil {
		log.Fatal(err)
	}
	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	return fileNames, nil
}

func serveApplication() {
	// create new gin router
	router := gin.Default()

	// Enable CORS
	router.Use(middleware.AllowAllCORS())

	router.GET("/", healthCheckHandler)

	router.GET("/files", getFileListHandler)

	router.Static("/assets", "./assets")

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
