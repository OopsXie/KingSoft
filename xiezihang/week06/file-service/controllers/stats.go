package controllers

import (
	"file-service/db"
	"file-service/models"
	"file-service/utils"

	"github.com/gin-gonic/gin"
)

func Stats(c *gin.Context) {
	type Result struct {
		Type      string
		Count     int64
		TotalSize int64
	}

	var results []Result
	db.DB.Model(&models.FileMeta{}).
		Select("type, count(*) as count, sum(size) as total_size").
		Group("type").Scan(&results)

	utils.Success(c, results)
}
