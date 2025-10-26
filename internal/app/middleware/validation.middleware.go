package middleware

import (
	"be-learn/internal/app/dto"
	"be-learn/internal/constants"
	"be-learn/internal/validatorx"
	"be-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Validate[T any](validateType constants.ValidateType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		
		validatekey, isErr := utils.GetTypeValidate(validateType)
		if isErr {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": []dto.ValidationError{{
					Field:   "type",
					Message: "Invalid validate type",
				}},
			})
			ctx.Abort()
			return
		}

		var bindErr error
		switch validatekey {
		case string(constants.ValidateBody):
			bindErr = ctx.ShouldBindJSON(&data)
		case string(constants.ValidateQuery):
			bindErr = ctx.ShouldBindQuery(&data)
		case string(constants.ValidateParam):
			bindErr = ctx.ShouldBindUri(&data)
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": []dto.ValidationError{{
					Field:   "type",
					Message: "Unknown validate type",
				}},
			})
			ctx.Abort()
			return
		}

		if bindErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": []dto.ValidationError{{
					Field:   validatekey,
					Message: bindErr.Error(),
				}},
			})
			ctx.Abort()
			return
		}


		if err := validatorx.Validate.Struct(data); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": utils.FormatValidationErrors(err),
			})
			ctx.Abort()
			return
		}

		ctx.Set(string(validateType), data)
		ctx.Next()
	}
}
