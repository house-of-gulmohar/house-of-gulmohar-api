package query

import (
	"house-of-gulmohar/internal/api/product/dto"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type ProductQuery struct{}

func (p *ProductQuery) GetAllProductsQuery(params dto.GetAllProductsDto) (string, []interface{}, error) {
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
			"b.image_url",
			"c.id",
			"c.name",
			"c.image_url",
		).
		From("product as p").
		InnerJoin("brand as b on p.brand = b.id").
		InnerJoin("category as c on p.category = c.id").
		Where("p.active = true")
	if params.Limit > 0 {
		psql = psql.
			Limit(uint64(params.Limit))
	}
	if params.Offset > 0 {
		psql = psql.Offset(uint64(params.Offset))
	}
	if len(params.Category) > 0 {
		psql = psql.Where(squirrel.Eq{"c.name": params.Category})
	}
	if len(params.Brand) > 0 {
		psql = psql.Where(squirrel.Eq{"b.name": params.Brand})
	}

	psql = psql.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.ToSql()

	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func (p *ProductQuery) GetProductQuery(params dto.GetProductDto) (string, error) {
	psql := squirrel.Select(
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
		Where("p.active = true").
		Where(squirrel.Eq{"p.id": params.Id}).PlaceholderFormat(squirrel.Dollar)

	query, _, err := psql.ToSql()
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return query, nil
}
