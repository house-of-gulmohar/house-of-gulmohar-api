package category

import (
	"context"
	"house-of-gulmohar/internal/api/category/dto"
	"house-of-gulmohar/internal/e"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/utils"
	"net/http"

	"github.com/jackc/pgx/v4"
)

type CategoryService struct {
	CategoryRepo CategoryRepo
}

func (s *CategoryService) GetAllCategories(ctx context.Context, params dto.GetAllCategoriesDto) (*model.Response, *model.ErrorResponse) {
	categories, count, err := s.CategoryRepo.GetAllCategories(ctx, params)
	if err != nil {
		return nil, e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "categories collected successfully",
		Data:    categories,
		Count:   count,
	}
	return &res, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, params dto.GetCategoryDto) (*model.Response, *model.ErrorResponse) {
	if len(params.Id) == 0 {
		return nil, e.ErrorMissingPathParam("id")
	}
	if err := utils.ValidateUUID(params.Id); err != nil {
		return nil, e.ErrorInvalidPathParam("id")
	}

	category, err := s.CategoryRepo.GetCategory(ctx, params)
	if err == pgx.ErrNoRows {
		return nil, e.NotFound("category not found")
	}
	if err != nil {
		return nil, e.InternalServerError(err.Error())
	}
	res := model.Response{
		Code:    http.StatusOK,
		Message: "category collected successfully",
		Data:    category,
	}
	return &res, nil
}
