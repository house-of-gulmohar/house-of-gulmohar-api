package app

import (
	"context"
	"encoding/json"
	"house-of-gulmohar/internal/data"
	"house-of-gulmohar/internal/e"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"
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

func (p *ProductHandler) handleGetProduct(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		return e.ErrorMissingPathParam("id")
	}
	if err := utils.ValidateUUID(id); err != nil {
		return e.ErrorInvalidPathParam("id")
	}
	product, err := p.ProductRepo.GetProduct(ctx, id)
	if err == pgx.ErrNoRows {
		return e.NotFound("product not found")
	}
	if err != nil {
		return e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "product collected successfully",
		Data:    product,
	}

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res.Send())
	return nil
}

func (p *ProductHandler) handleGetAllProductsByCategory(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
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
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		return e.ErrorMissingPathParam("id")
	}
	if err := utils.ValidateUUID(id); err != nil {
		return e.ErrorInvalidPathParam("id")
	}
	products, count, err := p.ProductRepo.GetAllProductsByCategory(ctx, id, limit, offset)
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
