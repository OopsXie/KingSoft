package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	TongyiKey   string
	DeepseekKey string
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	TongyiKey = os.Getenv("TONGYI_API_KEY")
	DeepseekKey = os.Getenv("DEEPSEEK_API_KEY")
}
