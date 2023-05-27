package handlers

import (
	"localflix/server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFileListHandler(c *gin.Context) {
	files := helper.GetFiles("assets")
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
