package telemetry

import (
	"context"
	"device-manager/internal/entity"
	"errors"
	"fmt"
	"time"
)

type Service struct {
	storage Storage
	databus DatabusProducer
}

var (
	ErrDeviceNotActive = errors.New("device not found or inactive")
)

func NewService(storage Storage, databus DatabusProducer) *Service {
	return &Service{
		storage: storage,
		databus: databus,
	}
}

func (s Service) GetTelemetryHistory(ctx context.Context, deviceID string, from time.Time, to time.Time) ([]entity.TelemetryData, error) {
	data, err := s.storage.GetHistoricalTelemetry(ctx, deviceID, from, to)
	if err != nil {
		return nil, fmt.Errorf("storage failed: %w", err)
	}

	return data, nil
}

func (s Service) GetTelemetryLatest(ctx context.Context, deviceID string) (*entity.TelemetryData, error) {
	data, err := s.storage.GetLatestTelemetry(ctx, deviceID)
	if err != nil {
		return nil, fmt.Errorf("storage failed: %w", err)
	}

	return data, nil
}

func (s Service) AddTelemetry(ctx context.Context, data entity.TelemetryData) error {
	// Проверка состояния устройства
	device, err := s.storage.GetDeviceById(ctx, data.DeviceId)
	if err != nil {
		return fmt.Errorf("failed to get device: %w", err)
	}

	if device == nil || device.Status != "active" {
		return ErrDeviceNotActive
	}

	// Добавление телеметрии
	err = s.storage.AddTelemetry(ctx, data.DeviceId, data.DeviceType, data.CreatedAt, data.TelemetryData)
	if err != nil {
		return fmt.Errorf("failed to add telemetry: %w", err)
	}

	// publish telemetry to databus
	err = s.databus.PublishTelemetry(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to publish telemetry in databus: %w", err)
	}

	return nil
}
