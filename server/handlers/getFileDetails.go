package handlers

import (
	"localflix/server/helper"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFileDetailsHandler(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath := filepath.Join("assets", fileName)

	fileInfo, err := os.Stat(filePath)
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

	if fileInfo.IsDir() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file",
		})
		return
	}

	fileData := helper.FileInfoData{
		Name:        fileInfo.Name(),
		Size:        helper.ByteCountSI(fileInfo.Size()),
		Path:        "/" + fileName,
		ModTime:     fileInfo.ModTime().Format(time.RFC3339),
		ChecksumSHA: helper.CalculateSHA256Checksum(filePath),
	}

	c.JSON(http.StatusOK, fileData)
}
