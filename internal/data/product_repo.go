package data

import (
	"context"
	"house-of-gulmohar/internal/constants"
	"house-of-gulmohar/internal/model"

	sq "github.com/Masterminds/squirrel"
)

type ProductRepo interface {
	GetAllProducts(ctx context.Context, limit int, offset int) ([]model.Product, int, error)
}

func (p *ProductDb) GetAllProducts(ctx context.Context, limit int, offset int) ([]model.Product, int, error) {
	defer ctx.Done()
	products := []model.Product{}
	qb := sq.Select(
		"product.id",
		"product.name",
		"product.description",
		"product.available",
		"product.mrp",
		"product.price",
		"product.on_sale",
		"product.discount",
		"product.brand",
		"product.category",
		"product.images",
		"product.replacement_period",
		"product.replacement_type",
		"product.warranty_period",
		"product.warranty_type",
		"product.created_at",
		"product.updated_at",
		"product.quantity",
		"product.featured",
	).
		From(constants.TABLES["PRODUCT"])
	if limit > 0 {
		qb.
			Limit(uint64(limit))
	}
	if offset > 0 {
		qb.Offset(uint64(offset))
	}

	sql, _, err := qb.ToSql()
	if err != nil {
		return nil, 0, err
	}

	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Available,
			&product.MRP,
			&product.Price,
			&product.OnSale,
			&product.Discount,
			&product.Brand,
			&product.Category,
			&product.Images,
			&product.ReplacementPeriod,
			&product.ReplacementType,
			&product.WarrantyPeriod,
			&product.WarrantyType,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Quantity,
			&product.Featured,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}
	return products, len(products), nil
}
