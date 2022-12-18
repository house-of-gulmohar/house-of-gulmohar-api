package data

import (
	"context"
	"house-of-gulmohar/internal/model/dto"

	"github.com/sirupsen/logrus"
)

type CategoryRepo interface {
	GetAllCategories(ctx context.Context, limit int, offset int) ([]dto.CategoryDto, int, error)
}

func (c *CategoryDb) GetAllCategories(ctx context.Context, limit int, offset int) ([]dto.CategoryDto, int, error) {
	defer ctx.Done()
	categories := []dto.CategoryDto{}
	sql, err := c.Query.GetAllCategoriesQuery(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	rows, err := c.Pool.Query(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}
	for rows.Next() {
		category := dto.CategoryDto{}
		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.ImageUrl,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, 0, err
		}
		categories = append(categories, category)
	}
	return categories, len(categories), nil
}
