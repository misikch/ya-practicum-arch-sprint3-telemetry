package devices

import (
	"context"

	"github.com/segmentio/kafka-go"

	"device-manager/internal/entity"
)

type Storage interface {
	CreateDevice(ctx context.Context, device *entity.Device) error
	UpdateDevice(ctx context.Context, deviceId string, status string) error
	GetDeviceById(ctx context.Context, deviceId string) (*entity.Device, error)
}

type Logger interface {
	Errorw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type KafkaReader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

type Metrics interface {
	Increment(name string)
}
