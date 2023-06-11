package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func GetLibraryHandler(c *gin.Context) {
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	// add error handling for file not found
	data, err := ioutil.ReadFile(libraryLocation + "/localflix-library.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var library LibraryData
	err = yaml.Unmarshal([]byte(data), &library)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	result := LibraryData{
		Version: library.Version,
		Videos:  library.Videos,
	}

	c.JSON(http.StatusOK, result)
}
