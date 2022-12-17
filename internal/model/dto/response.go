package dto

type ResponseDto struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
