package model

import "database/sql"

type Category struct {
	Id          string
	Name        string
	Description sql.NullString
	ImageURL    sql.NullString
	CreatedAt   string
	UpdatedAt   string
}
