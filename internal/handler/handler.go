package handler

type Handler struct {
	logger           Log
	telemetryService TelemetryService
}

func NewHandler(s TelemetryService, logger Log) *Handler {
	return &Handler{
		telemetryService: s,
		logger:           logger,
	}
}
