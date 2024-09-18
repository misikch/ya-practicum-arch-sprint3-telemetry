package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"device-manager/internal/entity"
)

// CreateDevice создает новое устройство.
func (s *Storage) CreateDevice(ctx context.Context, device *entity.Device) error {
	collection := s.Database().Collection("telemetry_devices")

	device.CreatedAt = time.Now()       // Устанавливаем время создания
	device.UpdatedAt = device.CreatedAt // Устанавливаем время обновления

	_, err := collection.InsertOne(ctx, device)
	if err != nil {
		s.logger.Errorw("failed to create device", "device", device, "error", err)
		return err
	}

	return nil
}

// UpdateDevice обновляет статус существующего устройства по идентификатору.
func (s *Storage) UpdateDevice(ctx context.Context, deviceId string, status string) error {
	collection := s.Database().Collection("telemetry_devices")
	filter := bson.M{"device_id": deviceId}
	updateBson := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(), // Устанавливаем время обновления
		},
	}

	_, err := collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		s.logger.Errorw("failed to update device", "deviceId", deviceId, "error", err)
		return err
	}

	return nil
}

// GetDeviceById получает устройство по идентификатору.
func (s *Storage) GetDeviceById(ctx context.Context, deviceId string) (*entity.Device, error) {
	collection := s.Database().Collection("telemetry_devices")
	filter := bson.M{"device_id": deviceId}

	var device entity.Device
	err := collection.FindOne(ctx, filter).Decode(&device)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		s.logger.Errorw("failed to get device", "deviceId", deviceId, "error", err)
		return nil, err
	}

	return &device, nil
}
