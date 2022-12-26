package data

import (
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BrandDb struct {
	Pool *pgxpool.Pool
	Db   *sql.DB
}
