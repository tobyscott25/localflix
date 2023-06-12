package handlers

import (
	"fmt"
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateFileInfoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateVideoDetailsHandler(c *gin.Context) {

	// // Check if the request body is empty - NOT WORKING
	// if c.Request.Body == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Empty request body"})
	// 	return
	// }
	// fmt.Println(c.Request.Body)

	var req UpdateFileInfoRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFileInfo := helper.FileInfoData{
		Title:       req.Title,
		Description: req.Description,
	}

	id := c.Param("id")

	library, err := helper.LoadLibraryFromYamlFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Missing Library file",
		})
		return
	}

	error := helper.UpdateVideoDetailsByID(library, id, updatedFileInfo)
	if error != nil {
		fmt.Println("Failed to update video details:", error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update video details. " + error.Error(),
		})
		return
	}

	videoDetails := helper.GetVideoDetailsByID(*library, id)

	helper.WriteLibraryToYamlFile(*library)

	c.JSON(http.StatusOK, videoDetails)
}
