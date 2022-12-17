package dto

type ResponseDto struct {
	Message string      `json:"msg"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponseDto struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	Detail  string      `json:"detail,omitempty"`
}
