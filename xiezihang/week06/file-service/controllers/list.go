package controllers

import (
	"file-service/db"
	"file-service/models"
	"file-service/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListFiles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	fileType := c.Query("type")

	offset := (page - 1) * pageSize
	var files []models.FileMeta
	query := db.DB.Order("upload_time desc").Offset(offset).Limit(pageSize)
	if fileType != "" {
		query = query.Where("type = ?", fileType)
	}
	query.Find(&files)

	utils.Success(c, files)
}
