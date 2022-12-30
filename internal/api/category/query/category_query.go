package query

import (
	"house-of-gulmohar/internal/api/category/dto"

	"github.com/Masterminds/squirrel"
)

type CategoryQuery struct{}

func (c *CategoryQuery) GetAllCategoriesQuery(params dto.GetAllCategoriesDto) (string, []interface{}, error) {
	psql := squirrel.
		Select(
			"c.id",
			"c.name",
			"c.description",
			"c.image_url",
			"c.created_at",
			"c.updated_at",
		).
		From("category as c")

	if params.Limit > 0 {
		psql = psql.
			Limit(uint64(params.Limit))
	}
	if params.Offset > 0 {
		psql = psql.Offset(uint64(params.Offset))
	}

	query, args, err := psql.ToSql()

	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func (c *CategoryQuery) GetCategoryQuery(params dto.GetCategoryDto) (string, []interface{}, error) {
	psql := squirrel.Select(
		"c.id",
		"c.name",
		"c.description",
		"c.image_url",
		"c.created_at",
		"c.updated_at",
	).
		From("category as c").
		Where(squirrel.Eq{"c.id": params.Id}).PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.ToSql()

	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}
