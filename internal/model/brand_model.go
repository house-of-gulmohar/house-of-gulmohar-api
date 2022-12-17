package model

import "database/sql"

type Brand struct {
	Id          string
	Name        string
	Description string
	ImageURL    sql.NullString
	CreatedAt   string
	UpdatedAt   string
}
