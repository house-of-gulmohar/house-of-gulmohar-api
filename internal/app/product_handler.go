package app

import (
	"house-of-gulmohar/internal/model"
	"net/http"
)

func (s *Server) getAllProducts(w http.ResponseWriter, r *http.Request) *model.ErrorResponse {
	return &model.ErrorResponse{Message: "hello world"}
}
