package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-clean-api/entity"
)

func DefineTagError(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("This field `%s` is required", field)
	case "email":
		return "Invalid email format"
	}
	return "Unknown error"
}

func BuildResponseError(err error) entity.ResponseError {
	var responseError entity.ResponseError

	//check type error from validator or not
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range errs {
			msg := DefineTagError(err.Tag(), err.Field())
			return entity.ResponseError{Message: &msg}
		}
	} else {
		msg := err.Error()
		return entity.ResponseError{Message: &msg}
	}

	return responseError
}
