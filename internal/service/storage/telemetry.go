package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"device-manager/internal/entity"
)

// GetLatestTelemetry получает последние данные телеметрии для указанного устройства.
func (s *Storage) GetLatestTelemetry(ctx context.Context, deviceId string) (*entity.TelemetryData, error) {
	collection := s.Database().Collection("telemetry")
	filter := bson.M{"deviceId": deviceId}
	opts := options.FindOne().SetSort(bson.D{
		{
			Key:   "createdAt",
			Value: -1,
		},
	})

	var telemetryData entity.TelemetryData
	err := collection.FindOne(ctx, filter, opts).Decode(&telemetryData)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		s.logger.Errorw("failed to get latest telemetry", "deviceId", deviceId, "error", err)
		return nil, err
	}

	return &telemetryData, nil
}

// GetHistoricalTelemetry получает исторические данные телеметрии для указанного устройства за определённый период времени.
func (s *Storage) GetHistoricalTelemetry(ctx context.Context, deviceId string, from, to time.Time) ([]entity.TelemetryData, error) {
	collection := s.Database().Collection("telemetry")
	filter := bson.M{
		"deviceId":  deviceId,
		"createdAt": bson.M{"$gte": from, "$lte": to},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		s.logger.Errorw("failed to get historical telemetry", "deviceId", deviceId, "error", err)
		return nil, err
	}
	// nolint:errcheck
	defer cursor.Close(ctx)

	var telemetryData []entity.TelemetryData
	if err = cursor.All(ctx, &telemetryData); err != nil {
		s.logger.Errorw("failed to decode historical telemetry", "deviceId", deviceId, "error", err)
		return nil, err
	}

	return telemetryData, nil
}

// AddTelemetry добавляет новые данные телеметрии для указанного устройства.
func (s *Storage) AddTelemetry(ctx context.Context, deviceId string, deviceType string, createdAt time.Time, telemetryData string) error {
	collection := s.Database().Collection("telemetry")
	data := entity.TelemetryData{
		DeviceId:      deviceId,
		DeviceType:    deviceType,
		CreatedAt:     createdAt,
		TelemetryData: telemetryData,
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		s.logger.Errorw("failed to insert telemetry data", "deviceId", deviceId, "error", err)
		return err
	}

	return nil
}
