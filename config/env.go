package config

import (
	"be-learn/utils"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT int `validate:"min=1,max=65535"`
}

var App Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Không tìm thấy file .env, dùng biến môi trường hệ thống")
	}

	App = Config{
		PORT: utils.GetEnv("PORT", 8000),
	}
}

