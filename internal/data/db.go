package data

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductDb struct {
	Db *pgxpool.Pool
}

type CategoryDb struct {
	Db *pgxpool.Pool
}

type BrandDb struct {
	Db *pgxpool.Pool
}
