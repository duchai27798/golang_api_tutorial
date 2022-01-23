package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Result  interface{} `json:"result"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, result interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Result:  result,
		Errors:  nil,
	}
}

func BuildErrorResponse(message string, error string, result interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Result:  result,
		Errors:  strings.Split(error, "\n"),
	}
}
