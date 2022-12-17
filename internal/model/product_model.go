package model

import "database/sql"

type Product struct {
	Id                string
	Name              string
	Description       sql.NullString
	Quantity          string
	Available         bool
	Featured          bool
	MRP               float64
	Price             float64
	OnSale            sql.NullBool
	Discount          float64
	Brand             string
	Category          string
	Images            []string
	ReplacementPeriod float64
	ReplacementType   string
	WarrantyPeriod    float64
	WarrantyType      string
	CreatedAt         string
	UpdatedAt         string
}
