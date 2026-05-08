package main

import (
	"file-service/db"
	"file-service/routes"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
