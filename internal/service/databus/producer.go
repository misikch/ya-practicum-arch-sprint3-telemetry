package databus

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"device-manager/internal/entity"
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

// PublishTelemetry публикует событие телеметрии в Kafka.
func (p *Producer) PublishTelemetry(ctx context.Context, msg entity.TelemetryData) error {
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
