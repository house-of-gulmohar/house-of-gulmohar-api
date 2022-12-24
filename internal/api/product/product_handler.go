package product

import (
	"context"
	"house-of-gulmohar/internal/api/product/dto"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type ProductHandler struct {
	*ProductService
}

func (p *ProductHandler) HandleGetAllProducts(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(req.URL.Query().Get("offset"))
	category := req.URL.Query().Get("category")
	brand := req.URL.Query().Get("brand")

	res, err := p.ProductService.GetAllProducts(ctx, dto.GetAllProductsDto{
		Limit:    limit,
		Offset:   offset,
		Category: category,
		Brand:    brand,
	})

	if err != nil {
		return err
	}

	utils.SendResponse(w, res)
	return nil
}

func (p *ProductHandler) HandleGetProduct(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	id := chi.URLParam(req, "id")

	res, err := p.ProductService.GetProduct(ctx, dto.GetProductDto{
		Id: id,
	})

	if err != nil {
		return err
	}

	utils.SendResponse(w, res)
	return nil
}
