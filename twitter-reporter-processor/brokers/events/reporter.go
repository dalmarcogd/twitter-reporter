package events

type reporterEvent struct {
	*EventBase
	ReporterId string `json:"reporter_id" validate:"required"`
	Tag        string `json:"tag" validate:"required"`
}

func NewReporterEvent(id string, tag string) *reporterEvent {
	return &reporterEvent{EventBase: NewEventBase("ReporterEvent"), ReporterId: id, Tag: tag}
}
