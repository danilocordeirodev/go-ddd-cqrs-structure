package query

import (
	"context"
)

type Query[Q any, R any] interface {
	Get(ctx context.Context, q Q) (R, error)
}

type Queries struct{}

func NewQueries(opts ...func(*Queries)) *Queries {
	q := &Queries{}
	for _, opt := range opts {
		opt(q)
	}
	return q
}
