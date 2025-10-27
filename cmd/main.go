package main

import (
	"be-learn/config"
	"be-learn/internal/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")

	router.RegisterAuthRoutes(api)
	

	r.Run(":" + strconv.Itoa(config.App.PORT))
}
