package controller

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetReadme(c *gin.Context) {

	cwd, _ := os.Getwd()

	readmePath := filepath.Join(cwd, "../README.md")

	content, err := os.ReadFile(readmePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "读取 README.md 失败",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "读取成功",
		"data": string(content),
	})
}
