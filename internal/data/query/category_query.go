package query

import "github.com/Masterminds/squirrel"

type CategoryQuery struct{}

func (c *CategoryQuery) GetAllCategoriesQuery(limit int, offset int) (string, error) {
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

	if limit > 0 {
		psql.
			Limit(uint64(limit))
	}
	if offset > 0 {
		psql.Offset(uint64(offset))
	}

	query, _, err := psql.ToSql()

	if err != nil {
		return "", err
	}
	return query, nil
}
