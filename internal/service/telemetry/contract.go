package telemetry

import (
	"context"
	"time"

	"device-manager/internal/service/databus"
	"device-manager/internal/service/storage"
)

type Storage interface {
	GetLatestTelemetry(ctx context.Context, deviceId string) (*storage.TelemetryData, error)
	GetHistoricalTelemetry(ctx context.Context, deviceId string, from, to time.Time) ([]storage.TelemetryData, error)
	AddTelemetry(ctx context.Context, deviceId string, deviceType string, createdAt time.Time, telemetryData map[string]interface{}) error
}

type DatabusProducer interface {
	PublishTelemetry(ctx context.Context, msg databus.TelemetryMessage) error
}
