package data

import (
	"context"
	"house-of-gulmohar/internal/data/query"
	"house-of-gulmohar/internal/model/dto"
)

type ProductRepo interface {
	GetAllProducts(ctx context.Context, limit int, offset int) ([]dto.ProductDto, int, error)
}

func (p *ProductDb) GetAllProducts(ctx context.Context, limit int, offset int) ([]dto.ProductDto, int, error) {
	defer ctx.Done()
	products := []dto.ProductDto{}
	sql, err := query.GetAllProductsQuery(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		product := dto.ProductDto{}
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.MRP,
			&product.Price,
			&product.OnSale,
			&product.Discount,
			&product.Images,
			&product.ReplacementPeriod,
			&product.ReplacementType,
			&product.WarrantyPeriod,
			&product.WarrantyType,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Quantity,
			&product.Featured,
			&product.Brand.Id,
			&product.Brand.Name,
			&product.Brand.ImageUrl,
			&product.Category.Id,
			&product.Category.Name,
			&product.Category.ImageUrl,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}
	return products, len(products), nil
}
