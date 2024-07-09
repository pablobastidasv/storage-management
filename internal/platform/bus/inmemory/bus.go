package inmemory

import (
	"context"
	"errors"
	"fmt"

	"co.bastriguez/inventory/kit/command"
)

type Bus struct {
	commandHandlers map[command.Type]command.CommandHandler
	queryHandlers   map[command.Type]command.QueryHandler
	eventHandlers   map[command.Type]command.EventHandler
}

func New() *Bus {
	return &Bus{
		commandHandlers: map[command.Type]command.CommandHandler{},
		queryHandlers:   map[command.Type]command.QueryHandler{},
		eventHandlers:   map[command.Type]command.EventHandler{},
	}
}

func (b *Bus) RegisterCommandHandler(t command.Type, handler command.CommandHandler) {
	b.commandHandlers[t] = handler
}

func (b *Bus) DispatchCommand(ctx context.Context, command command.Command) error {
	handler, ok := b.commandHandlers[command.Type()]
	if !ok {
		return errors.New("command not registered")
	}

	return handler.Handle(ctx, command)
}

func (b *Bus) RegisterQueryHandler(t command.Type, h command.QueryHandler) {
	b.queryHandlers[t] = h
}

func (b *Bus) DispatchQuery(ctx context.Context, query command.Query) (command.QueryResponse, error) {
	handler, ok := b.queryHandlers[query.Type()]
	if !ok {
		return nil, errors.New(fmt.Sprintf("query handler of type '%s' not registered", query.Type()))
	}

	return handler.Handle(ctx, query)
}

func (b *Bus) RegisterEventHandler(t command.Type, handler command.EventHandler) {
	b.eventHandlers[t] = handler
}

func (b *Bus) PublishEvent(ctx context.Context, event []command.Event) error {
	for _, e := range event {
		handler, ok := b.eventHandlers[e.Type()]
		if !ok {
			return fmt.Errorf("event handler for event type '%s' not registered", e.Type())
		}

        go func(evt command.Event) {
            err := handler.Handle(ctx, evt)
            if err != nil {
                // TODO: Do the logging of the error using zap (https://github.com/uber-go/zap)
            }
        }(e)
	}

	return nil
}

