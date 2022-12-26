package category

import (
	"context"
	"house-of-gulmohar/internal/api/category/dto"
	"house-of-gulmohar/internal/model"
	"house-of-gulmohar/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type CategoryHandler struct {
	*CategoryService
}

func (c *CategoryHandler) HandleGetAllCategories(w http.ResponseWriter, req *http.Request) *model.ErrorResponse {
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

	params := dto.GetAllCategoriesDto{
		Limit:  limit,
		Offset: offset,
	}

	res, errRes := c.CategoryService.GetAllCategories(ctx, params)
	if errRes != nil {
		return errRes
	}

	utils.SendResponse(w, res)
	return nil
}
