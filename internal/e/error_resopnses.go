package e

import (
	"fmt"
	"house-of-gulmohar/internal/model"
	"net/http"
)

func ErrorMissingQueryParam(p string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("missing %s in query", p),
	}
}

func ErrorInvalidQueryParam(p string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("missing %s in query", p),
	}
}

func ErrorMissingPathParam(p string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("missing %s in path", p),
	}
}

func ErrorInvalidPathParam(p string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("invalid %s in path", p),
	}
}

func ErrorInvalidField(f string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("invalid %s", f),
	}
}

func ErrorJsonParsing() *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "can't decode json",
	}
}

func InternalServerError(d string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "something happend, please try again later",
		Detail:  d,
	}
}

func NotFound(m string) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: m,
	}
}

func ValidationError(d interface{}) *model.ErrorResponse {
	return &model.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "invalid data provided, please verify",
		Data:    d,
	}
}
