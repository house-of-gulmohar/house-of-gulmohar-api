package e

import (
	"encoding/json"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/model/dto"
	"net/http"
)

type HandleException func(w http.ResponseWriter, r *http.Request) *model.Response

func (he HandleException) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := he(w, r)
	if err != nil {
		res := dto.ResponseDto{}
		res.Data = err.Data
		res.Message = err.Message
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(res)
	}
}
