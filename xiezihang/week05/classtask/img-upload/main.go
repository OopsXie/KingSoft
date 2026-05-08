package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	MaxUploadSize = 1 << 20 // 1MB
	UploadDir     = "./uploads"
)

func main() {
	r := gin.Default()

	// 创建上传目录
	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		panic("无法创建上传目录")
	}

	r.POST("/api/uploads", uploadHandler)
	r.GET("/api/preview/:filename", previewHandler)
	r.POST("/api/deleteimg", deleteHandler)

	r.Run(":8080")
}

func uploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析表单失败"})
		return
	}

	files := form.File["files"]
	var filenames []string

	for _, file := range files {
		if file.Size > MaxUploadSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("文件 %s 超过大小限制", file.Filename)})
			return
		}

		// 判断是否为图片
		if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("文件 %s 类型不是图片", file.Filename)})
			return
		}

		// 保存文件
		newName := uuid.New().String() + filepath.Ext(file.Filename)
		savePath := filepath.Join(UploadDir, newName)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
			return
		}

		filenames = append(filenames, newName)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "上传成功",
		"filenames": filenames,
	})
}

func previewHandler(c *gin.Context) {
	filename := c.Param("filename")
	fullPath := filepath.Join(UploadDir, filename)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
		return
	}

	if c.Query("download") == "1" {
		c.FileAttachment(fullPath, filename)
	} else {
		c.File(fullPath)
	}
}

func deleteHandler(c *gin.Context) {
	var req struct {
		Filename string `json:"filename"`
	}

	if err := c.BindJSON(&req); err != nil || req.Filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求无效"})
		return
	}

	fullPath := filepath.Join(UploadDir, req.Filename)
	if err := os.Remove(fullPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
