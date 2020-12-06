package helper

import (
	"gopkg.in/go-playground/validator.v9"
)

type ValidationRequest struct {
	Field string
	Tag   string
}

func (*ValidationRequest) ValidateHandling(model interface{}) []ValidationRequest {
	var errorResponse []ValidationRequest
	var validate = validator.New()
	isValid := validate.Struct(model)
	if isValid != nil {
		for _, err := range isValid.(validator.ValidationErrors) {
			errorResponse = append(errorResponse, ValidationRequest{err.Field(), err.Tag()})
		}

		return errorResponse
	}

	return nil
}
