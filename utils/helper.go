package utils

import (
	"be-learn/internal/app/dto"
	"be-learn/internal/constants"

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
