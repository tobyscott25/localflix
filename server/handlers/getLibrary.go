package handlers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func GetLibraryHandler(c *gin.Context) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	libraryFile, err := os.Open(libraryLocation + "/localflix-library.yaml")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Library file not found",
		})
		return
	}
	defer libraryFile.Close()

	libraryYamlData, err := io.ReadAll(libraryFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var library LibraryData
	err = yaml.Unmarshal(libraryYamlData, &library)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c.JSON(http.StatusOK, &library)
}
