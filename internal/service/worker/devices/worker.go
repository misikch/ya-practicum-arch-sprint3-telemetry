package devices

import (
	"context"
	"encoding/json"

	"device-manager/internal/entity"
)

type Worker struct {
	reader  KafkaReader
	logger  Logger
	storage Storage
	metrics Metrics
}

type DeviceMessage struct {
	DeviceId string `json:"deviceId"`
	Status   string `json:"status"`
}

const (
	metricDatabusFailed    = "worker.databus.failed"
	metricRepositoryFailed = "worker.repository.failed"
)

func NewWorker(reader KafkaReader, logger Logger, storage Storage, m Metrics) *Worker {
	return &Worker{
		reader:  reader,
		logger:  logger,
		storage: storage,
		metrics: m,
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		msg, err := w.reader.ReadMessage(ctx)
		if err != nil {
			w.logger.Errorw("failed to read message from kafka", "error", err)
			continue
		}

		var deviceMessage DeviceMessage
		err = json.Unmarshal(msg.Value, &deviceMessage)
		if err != nil {
			w.logger.Errorw("failed to unmarshal device message", "error", err)
			continue
		}

		w.logger.Infow("received device message from kafka", "deviceId", deviceMessage.DeviceId, "status", deviceMessage.Status)

		// Проверяем, существует ли устройство
		existingDevice, err := w.storage.GetDeviceById(ctx, deviceMessage.DeviceId)
		if err != nil {
			w.metrics.Increment(metricDatabusFailed)

			w.logger.Errorw("failed to get device by id", "deviceId", deviceMessage.DeviceId, "error", err)
			continue
		}

		if existingDevice == nil {
			// Устройство не существует, создаем новое
			device := &entity.Device{
				DeviceId: deviceMessage.DeviceId,
				Status:   deviceMessage.Status,
			}
			if err := w.storage.CreateDevice(ctx, device); err != nil {
				w.metrics.Increment(metricRepositoryFailed)

				w.logger.Errorw("failed to create device", "device", device, "error", err)
			}
		} else {
			// Устройство существует, обновляем статус
			if err := w.storage.UpdateDevice(ctx, deviceMessage.DeviceId, deviceMessage.Status); err != nil {
				w.metrics.Increment(metricRepositoryFailed)

				w.logger.Errorw("failed to update device", "deviceId", deviceMessage.DeviceId, "error", err)
			}
		}
	}
}
