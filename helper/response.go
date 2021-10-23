package helper

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse(status int, message string, data interface{}, errors interface{}) *BaseResponse {
	return &BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}
