package helpers

type Response struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Error:   false,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Error:   true,
		Data:    data,
	}

	return res
}
