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
		res := dto.ResponseDto{}
		res.Data = err.Data
		res.Message = err.Message
		if err.Code == 0 {
			err.Code = http.StatusOK
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
