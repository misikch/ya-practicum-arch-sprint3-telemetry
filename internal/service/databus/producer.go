package databus

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Producer struct {
	writer *kafka.Writer
	logger *zap.SugaredLogger
}

func NewProducer(writer *kafka.Writer, logger *zap.SugaredLogger) *Producer {
	return &Producer{
		writer: writer,
		logger: logger,
	}
}

type TelemetryMessage struct {
	DeviceId      string                 `json:"deviceId"`
	DeviceType    string                 `json:"deviceType"`
	CreatedAt     time.Time              `json:"createdAt"`
	TelemetryData map[string]interface{} `json:"telemetryData"`
}

// PublishTelemetry публикует событие телеметрии в Kafka.
func (p *Producer) PublishTelemetry(ctx context.Context, msg TelemetryMessage) error {
	message, err := json.Marshal(msg)
	if err != nil {
		p.logger.Errorw("failed to marshal telemetry message", "error", err)
		return err
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
	if err != nil {
		p.logger.Errorw("failed to send telemetry message to kafka", "error", err)
		return err
	}

	p.logger.Infow("successfully sent telemetry message to kafka", "deviceId", msg.DeviceId)
	return nil
}
