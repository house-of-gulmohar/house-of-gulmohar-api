package model

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

type ErrorResponse struct {
	Code    int
	Message string
	Data    interface{}
}
