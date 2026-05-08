package main

import (
	"homework/config"
	"homework/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()
	r.POST("/api/questions/create", handler.CreateQuestion)
	r.Run(":8081")
}
