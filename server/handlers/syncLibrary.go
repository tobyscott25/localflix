package handlers

import (
	"fmt"
	"localflix/server/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LibraryData struct {
	Version string                `json:"version"`
	Videos  []helper.FileInfoData `json:"videos"` // Currently only videos are supported
}

func SyncLibraryHandler(c *gin.Context) {

	version := os.Getenv("localflixSemanticVersion")
	libraryLocation := os.Getenv("LIBRARY_LOCATION")

	videosList := helper.GetAllVideosInDirectory(libraryLocation)

	data := LibraryData{
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
