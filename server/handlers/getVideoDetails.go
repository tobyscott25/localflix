package handlers

import (
	"fmt"
	"localflix/server/helper"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetVideoDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("ID:", id)

	filePath, err := helper.LookupFileByChecksum(id)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "File not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if fileInfo.IsDir() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file",
		})
		return
	}

	fileData := helper.FileInfoData{
		Name:           fileInfo.Name(),
		Size:           helper.HumanReadableFileSize(fileInfo.Size()),
		Path:           "/" + fileInfo.Name(),
		LastModified:   fileInfo.ModTime().Format(time.RFC3339),
		ChecksumSHA256: helper.CalculateSHA256Checksum(filePath),
	}

	c.JSON(http.StatusOK, fileData)
}
