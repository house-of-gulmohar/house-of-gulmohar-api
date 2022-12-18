package app

import (
	"context"
	"encoding/json"
	"house-of-gulmohar/internal/data"
	"house-of-gulmohar/internal/e"
	"house-of-gulmohar/internal/model"
	"net/http"
	"strconv"
	"time"
)

type ProductHandler struct {
	ProductRepo data.ProductRepo
}

func (p *ProductHandler) handleGetAllProducts(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(req.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	products, count, err := p.ProductRepo.GetAllProducts(ctx, limit, offset)
	if err != nil {
		return e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "products collected successfully",
		Data:    products,
		Count:   count,
	}

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res.Send())
	return nil
}
