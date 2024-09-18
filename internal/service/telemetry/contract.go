package telemetry

//go:generate mockgen -source=contract.go -destination=mock_contract.go -package=telemetry

import (
	"context"
	"time"

	"device-manager/internal/entity"
)

type Storage interface {
	GetLatestTelemetry(ctx context.Context, deviceId string) (*entity.TelemetryData, error)
	GetHistoricalTelemetry(ctx context.Context, deviceId string, from, to time.Time) ([]entity.TelemetryData, error)
	AddTelemetry(ctx context.Context, deviceId string, deviceType string, createdAt time.Time, telemetryData string) error
	GetDeviceById(ctx context.Context, deviceId string) (*entity.Device, error)
}

type DatabusProducer interface {
	PublishTelemetry(ctx context.Context, msg entity.TelemetryData) error
}
