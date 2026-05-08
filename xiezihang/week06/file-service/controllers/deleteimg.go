package controllers

import (
	"file-service/db"
	"file-service/models"
	"file-service/utils"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DeleteFile(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		utils.Fail(c, "缺少参数 filename")
		return
	}

	var file models.FileMeta
	if err := db.DB.First(&file, "id = ?", filename).Error; err != nil {
		utils.Fail(c, "文件不存在")
		return
	}

	filePath := filepath.Join("storage", file.Type, filename)
	if err := os.Remove(filePath); err != nil {
		utils.Fail(c, "删除文件失败")
		return
	}

	db.DB.Delete(&file)
	utils.Success(c, "删除成功")
}
