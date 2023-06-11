package handlers

import (
	"fmt"
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SyncLibraryHandler(c *gin.Context) {

	version := os.Getenv("localflixSemanticVersion")
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	videosList := helper.GetAllVideosInDirectory(libraryLocation)

	data := helper.LibraryData{
		Version: version,
		Videos:  videosList,
	}

	err := helper.WriteYAMLFile(libraryLocation+"/localflix-library.yaml", data)
	if err != nil {
		fmt.Printf("Failed to write YAML file: %v", err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
