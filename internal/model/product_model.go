package model

import (
	"house-of-gulmohar/internal/utils/types"
	"time"
)

type Product struct {
	Id                string           `json:"id"`
	Name              string           `json:"name"`
	Description       types.NullString `json:"description"`
	Quantity          int64            `json:"quantity"`
	Active            bool             `json:"-"`
	Featured          bool             `json:"featured"`
	MRP               float64          `json:"mrp"`
	Price             float64          `json:"price"`
	OnSale            bool             `json:"on_sale"`
	Discount          float64          `json:"discount"`
	Brand             string           `json:"brand"`
	Category          string           `json:"category"`
	Images            []string         `json:"images"`
	ReplacementPeriod int64            `json:"replacement_period"`
	ReplacementType   string           `json:"replacement_type"`
	WarrantyPeriod    int64            `json:"warranty_period"`
	WarrantyType      string           `json:"warranty_type"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}
