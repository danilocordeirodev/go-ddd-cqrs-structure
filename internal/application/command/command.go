package command

import (
	"context"
)

type Command[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type Commands struct{}

func NewCommands(opts ...func(*Commands)) *Commands {
	c := &Commands{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
