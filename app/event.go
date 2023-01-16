package app

import (
	"context"

	"github.com/dnozdrin/boilerplate/domain"
)

type EventHandler interface {
	Handle(ctx context.Context, event domain.Event)
}

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func (e *EventPublisher) Subscribe(handler EventHandler, events ...domain.Event) {
	for i := range events {
		e.handlers[events[i].Name()] = append(e.handlers[events[i].Name()], handler)
	}
}

func (e *EventPublisher) Notify(ctx context.Context, event domain.Event) {
	if event.IsAsynchronous() {
		go e.notify(ctx, event)
	}

	e.notify(ctx, event)
}

func (e *EventPublisher) notify(ctx context.Context, event domain.Event) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Handle(ctx, event)
	}
}
