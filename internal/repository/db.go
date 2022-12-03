package repository

import (
	"context"
	"fmt"

	vectorcalculator "github.com/AnimemeMops/hackathonProject/internal/vector-calculator"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func NewRepo(connStr string) (*Repo, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}

	return &Repo{
		conn: conn,
	}, nil
}

func (rp *Repo) Close(ctx context.Context) error {
	return rp.conn.Close(ctx)
}

func (rp *Repo) SaveVectors(ctx context.Context, tableName string, vectors []*vectorcalculator.IV) error {
	qb := sq.Insert(tableName).
	Columns("path", "vector").
	PlaceholderFormat(sq.Dollar)
	for _, iv := range vectors {
		qb.Values(iv.Path, iv.Vector)
	}

	sql, params, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	_, err = rp.conn.Exec(ctx, sql, params...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (rp *Repo) Vectors(ctx context.Context, tableName string, limit, offset uint64) ([]*vectorcalculator.IV, error) {
	qb := sq.Select(
		"path",
		"vector",
	).
	PlaceholderFormat(sq.Dollar).
	From(tableName)
	if limit != 0 {
		qb = qb.Limit(limit)
	}
	if offset != 0 {
		qb = qb.Offset(offset)
	}

	sql, params, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	rows, err := rp.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}
	defer rows.Close()

	var rowsData []*queryRow
	err = rows.Scan(&rowsData)
	if err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return convert(rowsData), nil
}

func convert(in []*queryRow) []*vectorcalculator.IV {
	out := make([]*vectorcalculator.IV, 0, len(in))

	for _, data := range in {
		out = append(out, &vectorcalculator.IV{
			Path: data.path,
			Vector: data.vector,
		})
	}

	return out
}
