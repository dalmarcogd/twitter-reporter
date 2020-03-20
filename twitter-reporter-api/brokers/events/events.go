package events

type reporterEvent struct {
	*EventBase
	ReporterId string `json:"account_id" validate:"required"`
	Tag        string `json:"tag" validate:"required"`
}

func (a reporterEvent) GetChannel() string {
	return a.Name
}

func NewReporterEvent(id string, tag string) *reporterEvent {
	return &reporterEvent{EventBase: NewEventBase("ReporterEvent"), ReporterId: id, Tag: tag}
}
