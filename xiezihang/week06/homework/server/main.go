package main

import (
	"log"
	"server/database"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("未加载 .env 文件，尝试使用系统环境变量")
	}

	// 初始化数据库
	database.InitDB()

	// 启动 Gin 路由
	r := router.SetupRouter()
	r.Run(":8080")
}
