package repository

import (
	"github.com/jackc/pgx/v4"
)

type Repo struct {
	conn *pgx.Conn
}

type queryRow struct {
	path string  `db:"path"`
	vector []float64  `db:"vector"`
}