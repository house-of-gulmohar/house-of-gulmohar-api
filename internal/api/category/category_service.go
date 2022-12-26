package category

import (
	"context"
	"house-of-gulmohar/internal/api/category/dto"
	"house-of-gulmohar/internal/e"
	"house-of-gulmohar/internal/model"
	"net/http"
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
