package controllers

import (
	"file-service/db"
	"file-service/models"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")

	var file models.FileMeta
	if err := db.DB.First(&file, "id = ?", filename).Error; err != nil {
		c.JSON(404, gin.H{"msg": "文件不存在"})
		return
	}

	filePath := filepath.Join("storage", file.Type, filename)
	c.FileAttachment(filePath, file.Filename)
}

func PreviewFile(c *gin.Context) {
	filename := c.Param("filename")

	var file models.FileMeta
	if err := db.DB.First(&file, "id = ?", filename).Error; err != nil {
		c.JSON(404, gin.H{"msg": "文件不存在"})
		return
	}

	filePath := filepath.Join("storage", file.Type, filename)
	c.File(filePath)
}
