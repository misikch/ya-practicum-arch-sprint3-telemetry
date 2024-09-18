package handler

import (
	"context"
	"time"

	"device-manager/internal/entity"
)

type TelemetryService interface {
	GetTelemetryLatest(ctx context.Context, deviceID string) (*entity.TelemetryData, error)
	GetTelemetryHistory(ctx context.Context, deviceID string, from time.Time, to time.Time) ([]entity.TelemetryData, error)
	AddTelemetry(ctx context.Context, data entity.TelemetryData) error
}

type Log interface {
	Error(args ...interface{})
}
