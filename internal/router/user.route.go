package router

import (
	"be-learn/config"
	"be-learn/internal/app/dto"
	"be-learn/internal/app/middleware"
	"be-learn/internal/app/model"
	"be-learn/internal/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", func(c *gin.Context) {
		var users []model.User
		config.DB.Where("deleted_at is null").Find(&users)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Lấy danh sách người dùng thành công",
			"data":    users,
		})
	})

	user.GET("/:UserID", middleware.Validate[dto.GetParamUser](constants.ValidateParam), func(c *gin.Context) {
		data, exists := c.Get("param")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Không tìm thấy body đã validate",
			})
			return
		}
		param := data.(dto.GetParamUser)
		var u model.User
		config.DB.Where("deleted_at is null").First(&u, param.UserID)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Lấy chi tiết người dùng thành công",
			"data":    u,
		})
	})

	user.POST("/", middleware.Validate[dto.CreateUserBody](constants.ValidateBody), func(c *gin.Context) {
		data, exists := c.Get("body")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Không tìm thấy body đã validate",
			})
			return
		}
		body := data.(dto.CreateUserBody)

		err := config.DB.Take(&model.User{}, "email = ?", body.Email).Error

		if err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Email đã tồn tại",
				"error":    "CONFLICT",
			})
			return 
		}

		u := model.User{
			Name: body.Name,
			Password: body.Password,
			Age: body.Age,
			Email: body.Email,
		}
		
		config.DB.Create(&u)


		c.JSON(http.StatusCreated, gin.H{
			"message": "Tạo người dùng thành công",
			"data":    u,
		})
	})

	user.PATCH("/:UserID",  middleware.Validate[dto.GetParamUser](constants.ValidateParam), middleware.Validate[dto.UpdateUserBody](constants.ValidateBody), func(c *gin.Context) {
		bodyData, existBody := c.Get("body")
		paramData, existParam := c.Get("param")
		if !existBody || !existParam {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Không tìm thấy body đã validate",
			})
			return
		}
		body := bodyData.(dto.UpdateUserBody)
		param := paramData.(dto.GetParamUser)
		
		var found model.User

		err := config.DB.First(&found, param.UserID).Error

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Không tìm thấy người dùng",
				"error":    "NOT_FOUND",
			})
			return 
		}
		
		config.DB.Model(&found).Updates(body)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Cập nhật người dùng thành công",
			"data":    found,
		})
	})

	user.DELETE("/:UserID",  middleware.Validate[dto.GetParamUser](constants.ValidateParam), func(c *gin.Context) {
		paramData, existParam := c.Get("param")
		if !existParam {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Không tìm thấy body đã validate",
			})
			return
		}
		param := paramData.(dto.GetParamUser)
		
		var found model.User

		err := config.DB.First(&found, param.UserID).Error

		log.Println(err, param.UserID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Không tìm thấy người dùng",
				"error":    "NOT_FOUND",
			})
			return 
		}

		config.DB.Delete(&found)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Xóa người dùng thành công",
			"data":    found,
		})
	})
}