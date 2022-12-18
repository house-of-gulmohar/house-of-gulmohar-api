package data

import (
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductDb struct {
	Pool *pgxpool.Pool
	Db   *sql.DB
}

type CategoryDb struct {
	Pool *pgxpool.Pool
	Db   *sql.DB
}

type BrandDb struct {
	Pool *pgxpool.Pool
	Db   *sql.DB
}
