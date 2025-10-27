package main

import (
	"be-learn/config"
	"be-learn/internal/app/middleware"
	"be-learn/internal/app/model"
	"be-learn/internal/router"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	err := config.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Lỗi migrate:", err)
	}

	m := config.DB.Migrator()
	m.CreateIndex(&model.User{}, "email_not_deleted")
	config.DB.Exec(`
			CREATE UNIQUE INDEX IF NOT EXISTS users_email_unique_not_deleted
			ON users (email)
			WHERE deleted_at IS NULL;
	`)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	api := r.Group("/api")

	router.RegisterAuthRoutes(api)

	router.RegisterUserRoutes(api)
	

	r.Run(":" + strconv.Itoa(config.App.PORT))
}
