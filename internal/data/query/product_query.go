package query

import "github.com/Masterminds/squirrel"

func GetAllProductsQuery(limit int, offset int) (string, error) {
	psql := squirrel.
		Select(
			"p.id",
			"p.name",
			"p.description",
			"p.mrp",
			"p.price",
			"p.on_sale",
			"p.discount",
			"p.images",
			"p.replacement_period",
			"p.replacement_type",
			"p.warranty_period",
			"p.warranty_type",
			"p.created_at",
			"p.updated_at",
			"p.quantity",
			"p.featured",
			"b.id",
			"b.name",
			"c.id",
			"c.name",
			"c.image_url",
			"b.image_url",
		).
		From("product as p").
		InnerJoin("brand as b on p.brand = b.id").
		InnerJoin("category as c on p.category = c.id").
		Where("p.active = true")
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