package shared

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DomainEvent define a interface que todos os eventos devem implementar
type DomainEvent interface {
	EventType() string
}

// BaseEvent contém os campos comuns a todos os eventos
type BaseEvent struct {
	ID        primitive.ObjectID
	AggID     primitive.ObjectID
	Timestamp time.Time
}

// NewBaseEvent cria um novo evento base
func NewBaseEvent(aggID primitive.ObjectID) BaseEvent {
	return BaseEvent{
		ID:        NewMongoID(),
		AggID:     aggID,
		Timestamp: time.Now(),
	}
}

// EventHandler define uma função que manipula eventos
type EventHandler func(context.Context, DomainEvent)

// EventBus gerencia a publicação e escuta de eventos de domínio
type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

// NewEventBus cria uma nova instância de EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

// Subscribe adiciona um handler para um tipo específico de evento
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

// Publish dispara um evento para os handlers registrados
func (eb *EventBus) Publish(ctx context.Context, event DomainEvent) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, ok := eb.handlers[event.EventType()]; ok {
		for _, handler := range handlers {
			go handler(ctx, event)
		}
	}
}
