package router

import (
	"be-learn/internal/app/dto"
	"be-learn/internal/app/middleware"
	"be-learn/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)


var user []dto.RegisterBody

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Lấy danh sách thành công",
			"data": user,
		})
	})
	
	auth.POST("/register", middleware.Validate[dto.RegisterBody](constants.ValidateBody), func(c *gin.Context) {
		data, exists := c.Get("body")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Không tìm thấy body đã validate",
        })
        return
    }

    u := data.(dto.RegisterBody)
		user = append(user, u)
		user = append(user, u)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Tạo người dùng thành công",
			"data": u,
		})
	})
}
