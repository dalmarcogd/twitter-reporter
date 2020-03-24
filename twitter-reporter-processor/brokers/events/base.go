package events

import (
	"context"
	"github.com/google/uuid"
	"go.elastic.co/apm"
	"time"
)

type Event interface {
	GetName() string
}

type EventTrace struct {
	SpanId string `json:"span_id" validate:"required"`
}

type EventBase struct {
	Uuid       string     `json:"uuid" validate:"required"`
	Timestamp  string     `json:"timestamp" validate:"required"`
	Name       string     `json:"name" validate:"required"`
	EventTrace EventTrace `json:"event_trace" validate:"required"`
}

func NewEventBase(ctx context.Context, name string) *EventBase {
	uid, _ := uuid.NewUUID()

	return &EventBase{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Uuid:      uid.String(),
		Name:      name,
		EventTrace: EventTrace{
			SpanId: apm.SpanFromContext(ctx).TraceContext().Span.String(),
		},
	}
}

func (e EventBase) GetName() string {
	return e.Name
}
