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

	// add error handling for file not found
	file, err := os.Open(libraryLocation + "/localflix-library.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var library LibraryData
	err = yaml.Unmarshal(data, &library)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	result := LibraryData{
		Version: library.Version,
		Videos:  library.Videos,
	}

	c.JSON(http.StatusOK, result)
}
