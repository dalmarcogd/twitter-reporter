package events

import "context"

type ReporterEvent struct {
	*EventBase
	ReporterId string `json:"reporter_id" validate:"required"`
	Tag        string `json:"tag" validate:"required"`
}

func NewReporterEvent(ctx context.Context, id string, tag string) *ReporterEvent {
	return &ReporterEvent{EventBase: NewEventBase(ctx, "ReporterEvent"), ReporterId: id, Tag: tag}
}
