package routes

import (
	"file-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8MB

	r.POST("/api/uploads", controllers.Upload)
	r.GET("/api/list", controllers.ListFiles)
	r.GET("/api/preview/:filename", controllers.PreviewFile)
	r.GET("/api/download/:filename", controllers.DownloadFile)
	r.GET("/api/stats", controllers.Stats)
	r.DELETE("/api/deleteimg", controllers.DeleteFile)

	return r
}
