package dto

import (
	"house-of-gulmohar/internal/utils/types"
	"time"
)

type CategoryDto struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Description types.NullString `json:"description"`
	ImageUrl    types.NullString `json:"image_url"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type GetAllCategoriesDto struct {
	Limit  int
	Offset int
}

type GetCategoryDto struct {
	Id string
}
