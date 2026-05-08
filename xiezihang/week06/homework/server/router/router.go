// package router

// import (
// 	"net/http"
// 	"server/controller"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRouter() *gin.Engine {
// 	r := gin.Default()

// 	// 静态文件托管（正式部署用）
// 	r.Static("/assets", "../client/dist/assets")
// 	r.LoadHTMLFiles("../client/dist/index.html")

// 	// 主页
// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 	})

// 	api := r.Group("/api")
// 	{
// 		api.GET("/questions", controller.GetQuestions)
// 		api.POST("/questions", controller.AddQuestion)
// 		api.PUT("/questions/:id", controller.UpdateQuestion)
// 		api.DELETE("/questions/:id", controller.DeleteOneQuestion) // 单个删除
// 		api.POST("/questions/delete", controller.DeleteQuestions)  // 批量删除

// 		api.POST("/questions/ai_generate", controller.GenerateByAI)
// 		api.GET("/readme", controller.GetReadme)
// 	}

// 	return r
// }

package router

import (
	"net/http"
	"server/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件托管（正式部署用）
	r.Static("/assets", "../client/dist/assets")
	r.LoadHTMLFiles("../client/dist/index.html")

	// 主页和问题页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/questions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/ai-generate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API路由组
	api := r.Group("/api")
	{
		api.GET("/questions", controller.GetQuestions)
		api.POST("/questions", controller.AddQuestion)
		api.PUT("/questions/:id", controller.UpdateQuestion)
		api.DELETE("/questions/:id", controller.DeleteOneQuestion) // 单个删除
		api.POST("/questions/delete", controller.DeleteQuestions)  // 批量删除

		api.POST("/questions/ai_generate", controller.GenerateByAI)
		api.GET("/readme", controller.GetReadme)
	}

	// 处理所有未匹配的路由，返回index.html
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return r
}
