package db

import (
	"file-service/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file_service.db"), &gorm.Config{}) // ✅ 使用 glebarez/sqlite
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.FileMeta{})
}
