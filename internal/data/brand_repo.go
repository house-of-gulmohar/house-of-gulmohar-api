package data

import "context"

type BrandRepo interface {
	GetAllBrands(ctx context.Context, limit int, offset int)
}

func (b *BrandDb) GetAllBrands(ctx context.Context, limit int, offset int) {}
