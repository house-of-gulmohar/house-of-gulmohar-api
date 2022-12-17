package data

import "context"

type CategoryRepo interface {
	GetAllCategories(ctx context.Context, limit int, offset int)
}

func (c *CategoryDb) GetAllCategories(ctx context.Context, limit int, offset int) {}
