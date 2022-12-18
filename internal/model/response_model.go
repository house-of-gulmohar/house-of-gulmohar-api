package model

import "house-of-gulmohar/internal/model/dto"

type Response struct {
	Code    int
	Message string
	Count   int
	Data    interface{}
}

type ErrorResponse struct {
	Code    int
	Message string
	Data    interface{}
	Detail  string
}

func (r *Response) Send() *dto.ResponseDto {
	return &dto.ResponseDto{
		Message: r.Message,
		Data:    r.Data,
		Count:   r.Count,
	}
}
