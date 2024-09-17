package handler

import (
	"context"
	"time"

	"device-manager/internal/service/storage"
)

type TelemetryService interface {
	GetTelemetryHistory(ctx context.Context, deviceID string, from time.Time, to time.Time) ([]storage.TelemetryData, error)
}

type Log interface {
	Error(args ...interface{})
}
