package product

import (
	"context"
	"house-of-gulmohar/internal/api/product/dto"
	"house-of-gulmohar/internal/e"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/utils"
	"net/http"

	"github.com/jackc/pgx/v4"
)

type ProductService struct {
	ProductRepo ProductRepo
}

func (s *ProductService) GetAllProducts(ctx context.Context, params dto.GetAllProductsDto) (*model.Response, *model.ErrorResponse) {
	products, count, err := s.ProductRepo.GetAllProducts(ctx, params)
	if err != nil {
		return nil, e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "products collected successfully",
		Data:    products,
		Count:   count,
	}
	return &res, nil
}

func (s *ProductService) GetProduct(ctx context.Context, params dto.GetProductDto) (*model.Response, *model.ErrorResponse) {
	if len(params.Id) == 0 {
		return nil, e.ErrorMissingPathParam("id")
	}
	if err := utils.ValidateUUID(params.Id); err != nil {
		return nil, e.ErrorInvalidPathParam("id")
	}
	product, err := s.ProductRepo.GetProduct(ctx, params)
	if err == pgx.ErrNoRows {
		return nil, e.NotFound("product not found")
	}
	if err != nil {
		return nil, e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "product collected successfully",
		Data:    product,
	}
	return &res, nil
}
