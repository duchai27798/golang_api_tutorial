package helper

// Response response struct
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []ApiError  `json:"errors"`
	Result  interface{} `json:"result"`
}

// EmptyObj empty object
type EmptyObj struct{}

// BuildResponse build successful response object
func BuildResponse(status bool, message string, result interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Result:  result,
		Errors:  nil,
	}
}

// BuildErrorResponse build error response object
func BuildErrorResponse(message string, error []ApiError, result interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Result:  result,
		Errors:  error,
	}
}
