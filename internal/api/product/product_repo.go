package product

import (
	"context"
	"database/sql"
	"house-of-gulmohar/internal/api/product/dto"
	"house-of-gulmohar/internal/api/product/query"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type ProductDb struct {
	Pool  *pgxpool.Pool
	Db    *sql.DB
	Query query.ProductQuery
}

type ProductRepo interface {
	GetAllProducts(ctx context.Context, params dto.GetAllProductsDto) ([]dto.ProductDto, int, error)
	GetProduct(ctx context.Context, params dto.GetProductDto) (*dto.ProductDto, error)
	GetAllProductsByCategory(ctx context.Context, categoryId string, limit int, offset int) ([]dto.ProductDto, int, error)
}

func (p *ProductDb) GetAllProducts(ctx context.Context, params dto.GetAllProductsDto) ([]dto.ProductDto, int, error) {
	defer ctx.Done()
	products := []dto.ProductDto{}
	sql, err := p.Query.GetAllProductsQuery(params)
	if err != nil {
		return nil, 0, err
	}
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		logrus.Error(err)
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
			logrus.Error(err)
			return nil, 0, err
		}
		products = append(products, product)
	}
	return products, len(products), nil
}

func (p *ProductDb) GetProduct(ctx context.Context, params dto.GetProductDto) (*dto.ProductDto, error) {
	defer ctx.Done()
	product := dto.ProductDto{}
	sql, err := p.Query.GetProductQuery(params)
	if err != nil {
		return nil, err
	}
	row := p.Pool.QueryRow(context.Background(), sql, params.Id)
	err = row.Scan(
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
		if err != pgx.ErrNoRows {
			logrus.Error(err)
		}
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) GetAllProductsByCategory(ctx context.Context, categoryId string, limit int, offset int) ([]dto.ProductDto, int, error) {
	defer ctx.Done()
	products := []dto.ProductDto{}
	sql, err := p.Query.GetAllProductsByCategoryQuery(categoryId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	rows, err := p.Pool.Query(ctx, sql, categoryId)
	if err != nil {
		logrus.Error(err)
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
