package config

import (
	"be-learn/utils"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT 		int 		`validate:"min=1,max=65535"`
	DB_USER string 
	DB_PASS string
	DB_HOST string
	DB_PORT int			`validate:"min=1,max=65535"`
	DB_NAME string
}

var App Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Không tìm thấy file .env, dùng biến môi trường hệ thống")
	}

	App = Config{
		PORT: utils.GetEnv("PORT", 8000),
		DB_USER: utils.GetEnv("DB_USER", "root"),
		DB_PASS: utils.GetEnv("DB_PASS", "root"),
		DB_HOST: utils.GetEnv("DB_HOST", "root"),
		DB_PORT: utils.GetEnv("DB_PORT", 5432),
		DB_NAME: utils.GetEnv("DB_NAME", "root"),
	}
}

