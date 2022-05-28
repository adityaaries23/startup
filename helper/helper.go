package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Metadata    `json:"meta"`
	Data interface{} `json:"data"`
}

type Metadata struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func FormatResponse(message string, code int, status string, Data interface{}) Response {
	Metadata := Metadata{
		Message: message,
		Code:    code,
		Status:  status,
	}
	Reponse := Response{
		Meta: Metadata,
		Data: Data,
	}
	return Reponse
}

func FormatErrorValidation(err error) []string {
	var errors []string
	for _, b := range err.(validator.ValidationErrors) {
		errors = append(errors, b.Error())
	}
	return errors
}
