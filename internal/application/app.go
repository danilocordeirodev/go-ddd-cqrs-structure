package application

import (
	"github.com/danilocordeirodev/go-ddd-cqrs-structure/internal/application/command"
	"github.com/danilocordeirodev/go-ddd-cqrs-structure/internal/application/query"
)

type App struct {
	Commands *command.Commands
	Queries  *query.Queries
}

func New(
	commands *command.Commands,
	queries *query.Queries,
) *App {
	return &App{
		Commands: commands,
		Queries:  queries,
	}
}
