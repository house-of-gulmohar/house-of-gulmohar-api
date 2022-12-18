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

type CategoryHandler struct {
	CategoryRepo data.CategoryRepo
}

func (c *CategoryHandler) handleGetAllCategories(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
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
	categories, count, err := c.CategoryRepo.GetAllCategories(ctx, limit, offset)
	if err != nil {
		return e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "categories collected successfully",
		Data:    categories,
		Count:   count,
	}

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res.Send())
	return nil
}

func (c *CategoryHandler) handleGetCategory(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		return e.ErrorMissingPathParam("id")
	}
	if err := utils.ValidateUUID(id); err != nil {
		return e.ErrorInvalidPathParam("id")
	}
	product, err := c.CategoryRepo.GetCategory(ctx, id)
	if err == pgx.ErrNoRows {
		return e.NotFound("category not found")
	}
	if err != nil {
		return e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "category collected successfully",
		Data:    product,
	}
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res.Send())
	return nil
}
