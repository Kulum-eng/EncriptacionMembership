package responses

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message, error string) Response {
	return Response{
		Message: message,
		Data:    error,
	}
}
