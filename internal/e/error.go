package e

import (
	"encoding/json"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/model/dto"
	"net/http"
)

type ExeptionHandler func(w http.ResponseWriter, r *http.Request) *model.ErrorResponse

func (eh ExeptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := eh(w, r)
	if err != nil {
		res := dto.ErrorResponseDto{}
		res.Data = err.Data
		res.Message = err.Message
		res.Detail = err.Detail
		if err.Code == 0 {
			err.Code = http.StatusBadRequest
		}
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(res)
	}
}

func HandleException(eh ExeptionHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eh.ServeHTTP(w, r)
	})
}
