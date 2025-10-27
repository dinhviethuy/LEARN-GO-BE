package utils

import (
	"be-learn/internal/app/dto"
	"be-learn/internal/constants"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) []dto.ValidationError {
	var errs []dto.ValidationError
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			errs = append(errs, dto.ValidationError{
				Field:   e.Field(),
				Message: getMessage(e),
			})
		}
	} else if err != nil {
		errs = append(errs, dto.ValidationError{
			Field:   "",
			Message: err.Error(),
		})
	}
	return errs
}


func getMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return e.Field() + " must be a valid email"
	case "min":
		return e.Field() + " must have at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must have at most " + e.Param() + " characters"
	case "gte":
		return e.Field() + " must be >= " + e.Param()
	case "lte":
		return e.Field() + " must be <= " + e.Param()
	default:
		return e.Field() + " is invalid"
	}
}

func GetTypeValidate(t constants.ValidateType) (string, bool) {
	switch t {
	case constants.ValidateBody:
		return "body", false
	case constants.ValidateQuery:
		return "query", false
	case constants.ValidateParam:
		return "param", false
	default:
		return "unknown", true
	}
}

func GetEnv[T any](key string, fallback T) T {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	var result any
	var err error

	switch any(fallback).(type) {
	case int:
		result, err = strconv.Atoi(value)
	case int64:
		var v int64
		v, err = strconv.ParseInt(value, 10, 64)
		result = v
	case float64:
		var v float64
		v, err = strconv.ParseFloat(value, 64)
		result = v
	case bool:
		var v bool
		v, err = strconv.ParseBool(value)
		result = v
	case string:
		result = value
	default:
		log.Fatalf("Không hỗ trợ kiểu dữ liệu cho key %s", key)
	}

	if err != nil {
		log.Printf("Không parse được %s=%q, dùng giá trị mặc định", key, value)
		return fallback
	}

	return result.(T)
}