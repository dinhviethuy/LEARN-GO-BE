package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bắt tất cả panic
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Lỗi hệ thống",
					"error":   r,
				})
			}
		}()

		c.Next() // tiếp tục request chain

		// Kiểm tra nếu có lỗi trong context
		errs := c.Errors
		if len(errs) > 0 {
			lastErr := errs.Last()

			// Có thể tùy chỉnh theo loại lỗi
			switch lastErr.Type {
			case gin.ErrorTypeBind:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Dữ liệu gửi lên không hợp lệ",
					"error":   lastErr.Error(),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Đã xảy ra lỗi",
					"error":   lastErr.Error(),
				})
			}
		}
	}
}
