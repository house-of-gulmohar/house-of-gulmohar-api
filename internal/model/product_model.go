package model

import (
	"house-of-gulmohar/internal/utils"
	"time"
)

type Product struct {
	Id                string
	Name              string
	Description       utils.NullString
	Quantity          int64
	Available         bool
	Featured          bool
	MRP               float64
	Price             float64
	OnSale            bool
	Discount          float64
	Brand             string
	Category          string
	Images            []string
	ReplacementPeriod int64
	ReplacementType   string
	WarrantyPeriod    int64
	WarrantyType      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
