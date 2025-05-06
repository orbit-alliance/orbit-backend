package shared

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DomainEvent defines the interface that all events must implement
type DomainEvent interface {
	EventType() string
}

// BaseEvent contains fields common to all events
type BaseEvent struct {
	ID        primitive.ObjectID
	Timestamp time.Time
}

func NewBaseEvent(aggID primitive.ObjectID) BaseEvent {
	return BaseEvent{
		ID:        NewMongoID(),
		Timestamp: time.Now(),
	}
}

// EventHandler defines a function that handles events
type EventHandler func(context.Context, DomainEvent)

// EventBus manages publishing and listening of domain events
type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

// Subscribe adds a handler for a specific type of event
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

// Publish triggers an event to registered handlers
func (eb *EventBus) Publish(ctx context.Context, event DomainEvent) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, ok := eb.handlers[event.EventType()]; ok {
		for _, handler := range handlers {
			go handler(ctx, event)
		}
	}
}
