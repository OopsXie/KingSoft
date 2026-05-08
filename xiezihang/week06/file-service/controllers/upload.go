package controllers

import (
	"file-service/db"
	"file-service/models"
	"file-service/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var allowedExts = map[string]bool{
	".jpg": true, ".png": true, ".js": true, ".css": true, ".html": true,
}

const maxFileSize = 5 << 20 // 5MB

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		utils.Fail(c, "无法读取上传文件")
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		utils.Fail(c, "没有文件被上传")
		return
	}

	var uploaded []string

	for _, file := range files {
		if file.Size > maxFileSize {
			utils.Fail(c, "文件过大："+file.Filename)
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename)) // e.g. .jpg
		if !allowedExts[ext] {
			utils.Fail(c, "不支持的文件类型："+file.Filename)
			return
		}

		fileType := ext[1:] // 去掉“.”
		targetDir := filepath.Join("storage", fileType)
		os.MkdirAll(targetDir, os.ModePerm)

		uid := uuid.New().String()
		dst := filepath.Join(targetDir, uid+ext)

		srcFile, err := file.Open()
		if err != nil {
			utils.Fail(c, "无法打开文件："+file.Filename)
			return
		}
		defer srcFile.Close()

		outFile, err := os.Create(dst)
		if err != nil {
			utils.Fail(c, "保存文件失败："+file.Filename)
			return
		}
		defer outFile.Close()

		io.Copy(outFile, srcFile)

		db.DB.Create(&models.FileMeta{
			ID:         uid + ext,
			Filename:   file.Filename,
			Size:       file.Size,
			Type:       fileType,
			UploadTime: time.Now().Unix(),
		})

		uploaded = append(uploaded, uid+ext)
	}

	utils.Success(c, gin.H{"uploaded": uploaded})
}
