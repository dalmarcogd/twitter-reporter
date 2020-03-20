package events

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	GetName() string
	GetChannel() string
}

type EventBase struct {
	Uuid      string `json:"uuid" validate:"required"`
	Timestamp string `json:"timestamp" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

func NewEventBase(name string) *EventBase {
	uid, _ := uuid.NewUUID()
	return &EventBase{Uuid: uid.String(), Timestamp: time.Now().UTC().Format(time.RFC3339), Name: name}
}

func (e EventBase) GetName() string {
	return e.Name
}
