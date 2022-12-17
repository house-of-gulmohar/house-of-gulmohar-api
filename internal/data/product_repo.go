package data

import "context"

type ProductRepo interface {
	GetAllProducts(ctx context.Context, limit int, offset int)
}

func (p *ProductDb) GetAllProducts(ctx context.Context, limit int, offset int) {}
