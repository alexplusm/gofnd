package gofnd

type responseBaseTemplate struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildSuccessResponse(data interface{}) interface{} {
	return &responseBaseTemplate{Success: true, Data: data}
}

func BuildFailureResponse(message string, data interface{}) interface{} {
	return &responseBaseTemplate{Success: false, Message: message, Data: data}
}
