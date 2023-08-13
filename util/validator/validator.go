package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	alphaSpaceRegexString string = "^[a-zA-Z ]+$"
	passwordString        string = `^[a-zA-Z0-9!@#\$%\^&\*]+$`
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("pwd", isPwd)
	validate.RegisterValidation("alpha_space", isAlphaSpace)

	return validate
}

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "email":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid email address", err.Field())
			case "min":
				resp.Errors[i] = fmt.Sprintf("%s must be a minimum of %s in length", err.Field(), err.Param())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "alpha_space":
				resp.Errors[i] = fmt.Sprintf("%s can only contain alphabetic and space characters", err.Field())
			case "pwd":
				resp.Errors[i] = fmt.Sprintf("%s invalid", err.Field())
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}

func isAlphaSpace(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(alphaSpaceRegexString)
	return reg.MatchString(fl.Field().String())
}

func isPwd(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(passwordString)
	return reg.MatchString(fl.Field().String())
}
