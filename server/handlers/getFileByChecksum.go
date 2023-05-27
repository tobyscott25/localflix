package handlers

import (
	"localflix/server/helper"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFileByChecksumHandler(c *gin.Context) {
	checksum := c.Param("checksum")

	filePath, err := helper.LookupFileByChecksum(checksum)
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
		Name:        fileInfo.Name(),
		Size:        helper.ByteCountSI(fileInfo.Size()),
		Path:        "/" + fileInfo.Name(),
		ModTime:     fileInfo.ModTime().Format(time.RFC3339),
		ChecksumSHA: helper.CalculateSHA256Checksum(filePath),
	}

	c.JSON(http.StatusOK, fileData)
}
