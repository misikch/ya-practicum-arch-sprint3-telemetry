package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Storage struct {
	client *mongo.Client
	logger *zap.SugaredLogger
	dbName string
}

// New создаёт новый экземпляр хранилища (Storage) с использованием MongoDB.
func New(client *mongo.Client, logger *zap.SugaredLogger, dbName string) *Storage {
	return &Storage{
		client: client,
		logger: logger,
		dbName: dbName,
	}
}

// Database предоставляет доступ к базе данных MongoDB.
func (s *Storage) Database() *mongo.Database {
	return s.client.Database(s.dbName)
}

type TelemetryData struct {
	DeviceId      string                 `bson:"deviceId" json:"deviceId"`
	DeviceType    string                 `bson:"deviceType" json:"deviceType"`
	CreatedAt     time.Time              `bson:"createdAt" json:"createdAt"`
	TelemetryData map[string]interface{} `bson:"telemetryData" json:"telemetryData"`
}

// GetLatestTelemetry получает последние данные телеметрии для указанного устройства.
func (s *Storage) GetLatestTelemetry(ctx context.Context, deviceId string) (*TelemetryData, error) {
	collection := s.Database().Collection("telemetry")
	filter := bson.M{"deviceId": deviceId}
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})

	var telemetryData TelemetryData
	err := collection.FindOne(ctx, filter, opts).Decode(&telemetryData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		s.logger.Errorw("failed to get latest telemetry", "deviceId", deviceId, "error", err)
		return nil, err
	}

	return &telemetryData, nil
}

// GetHistoricalTelemetry получает исторические данные телеметрии для указанного устройства за определённый период времени.
func (s *Storage) GetHistoricalTelemetry(ctx context.Context, deviceId string, from, to time.Time) ([]TelemetryData, error) {
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
	defer cursor.Close(ctx)

	var telemetryData []TelemetryData
	if err = cursor.All(ctx, &telemetryData); err != nil {
		s.logger.Errorw("failed to decode historical telemetry", "deviceId", deviceId, "error", err)
		return nil, err
	}

	return telemetryData, nil
}

// AddTelemetry добавляет новые данные телеметрии для указанного устройства.
func (s *Storage) AddTelemetry(ctx context.Context, deviceId string, deviceType string, createdAt time.Time, telemetryData map[string]interface{}) error {
	collection := s.Database().Collection("telemetry")
	data := TelemetryData{
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
