package command

import "golang.org/x/net/context"

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name Bus
type Bus interface {
	RegisterCommandHandler(Type, CommandHandler)
	DispatchCommand(context.Context, Command) error
	RegisterQueryHandler(Type, QueryHandler)
	DispatchQuery(context.Context, Query) (QueryResponse, error)
	RegisterEventHandler(Type, EventHandler)
	PublishEvent(context.Context, []Event) error
}

type Type string

type Command interface {
	Type() Type
}

type CommandHandler interface {
	Handle(context.Context, Command) error
}

type Query interface {
	Type() Type
}

type QueryResponse interface{}

type QueryHandler interface {
	Handle(context.Context, Query) (QueryResponse, error)
}

type Event interface {
	Type() Type
}

type EventHandler interface {
	Handle(context.Context, Event) error
}
