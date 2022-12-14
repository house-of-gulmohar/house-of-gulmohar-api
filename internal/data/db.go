package data

import (
	"database/sql"
	"house-of-gulmohar/internal/data/query"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CategoryDb struct {
	Pool  *pgxpool.Pool
	Db    *sql.DB
	Query query.CategoryQuery
}

type BrandDb struct {
	Pool *pgxpool.Pool
	Db   *sql.DB
}
