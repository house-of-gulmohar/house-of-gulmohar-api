package e

import (
	"fmt"
	"house-of-gulmohar/internal/model"
	"net/http"
)

func ErrorMissingQueryParam(p string) *model.Response {
	return &model.Response{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("missing %s in query", p),
	}
}

func ErrorMissingPathParam(p string) *model.Response {
	return &model.Response{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("missing %s in path", p),
	}
}

func ErrorInvalidField(f string) *model.Response {
	return &model.Response{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("invalid %s", f),
	}
}

func ErrorJsonParsing() *model.Response {
	return &model.Response{
		Code:    http.StatusInternalServerError,
		Message: "can't decode json",
	}
}

func InternalServerError() *model.Response {
	return &model.Response{
		Code:    http.StatusInternalServerError,
		Message: "something happend, please try again later",
	}
}

func NotFound(m string) *model.Response {
	return &model.Response{
		Code:    http.StatusInternalServerError,
		Message: m,
	}
}

func ValidationError(d interface{}) *model.Response {
	return &model.Response{
		Code:    http.StatusBadRequest,
		Message: "invalid data provided, please verify",
		Data:    d,
	}
}
