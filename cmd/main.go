package main

import (
	"be-learn/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")

	router.RegisterAuthRoutes(api)
	

	r.Run(":8000")
}
