package telemetry

import (
	"context"
	"fmt"
	"time"

	"device-manager/internal/service/storage"
)

type Service struct {
	storage Storage
	databus DatabusProducer
}

func NewService(storage Storage, databus DatabusProducer) *Service {
	return &Service{
		storage: storage,
		databus: databus,
	}
}

func (s Service) GetTelemetryHistory(ctx context.Context, deviceID string, from time.Time, to time.Time) ([]storage.TelemetryData, error) {
	data, err := s.storage.GetHistoricalTelemetry(ctx, deviceID, from, to)
	if err != nil {
		return nil, fmt.Errorf("storage failed: %w", err)
	}

	return data, nil
}
