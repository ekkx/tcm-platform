package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Execer interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

func New(execer Execer) *Queries {
	return &Queries{
		execer: execer,
	}
}

type Queries struct {
	execer Execer
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		execer: tx,
	}
}
