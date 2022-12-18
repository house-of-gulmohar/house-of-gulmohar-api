package dto

import (
	"house-of-gulmohar/internal/utils"
	"time"
)

type CategoryDto struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Description utils.NullString `json:"description"`
	ImageUrl    utils.NullString `json:"image_url"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
