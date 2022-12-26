package category

import (
	"context"
	"database/sql"
	"house-of-gulmohar/internal/api/category/dto"
	"house-of-gulmohar/internal/api/category/query"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type CategoryDb struct {
	Pool  *pgxpool.Pool
	Db    *sql.DB
	Query query.CategoryQuery
}

type CategoryRepo interface {
	GetAllCategories(ctx context.Context, params dto.GetAllCategoriesDto) ([]dto.CategoryDto, int, error)
	GetCategory(ctx context.Context, id string) (*dto.CategoryDto, error)
}

func (c *CategoryDb) GetAllCategories(ctx context.Context, params dto.GetAllCategoriesDto) ([]dto.CategoryDto, int, error) {
	defer ctx.Done()
	categories := []dto.CategoryDto{}
	sql, args, err := c.Query.GetAllCategoriesQuery(params)
	if err != nil {
		return nil, 0, err
	}
	rows, err := c.Pool.Query(ctx, sql, args...)
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

func (c *CategoryDb) GetCategory(ctx context.Context, id string) (*dto.CategoryDto, error) {
	defer ctx.Done()
	category := dto.CategoryDto{}
	sql, args, err := c.Query.GetCategoryQuery(id)
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.ImageUrl,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err != pgx.ErrNoRows {
			logrus.Error(err)
		}
		return nil, err
	}
	return &category, nil
}
