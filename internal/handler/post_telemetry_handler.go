package handler

import (
	"context"
	"errors"

	"device-manager/api"
	"device-manager/internal/entity"
	"device-manager/internal/service/telemetry"
)

// TODO: process batch requests
func (h Handler) DevicesDeviceIDTelemetryPost(
	ctx context.Context,
	req *api.DevicesDeviceIDTelemetryPostReq,
	params api.DevicesDeviceIDTelemetryPostParams,
) (api.DevicesDeviceIDTelemetryPostRes, error) {
	td := entity.TelemetryData{
		DeviceId:      params.DeviceID,
		DeviceType:    req.DeviceType,
		CreatedAt:     req.CreatedAt,
		TelemetryData: req.TelemetryData,
	}

	err := h.telemetryService.AddTelemetry(ctx, td)

	if err != nil {
		if errors.Is(err, telemetry.ErrDeviceNotActive) {
			return &api.DevicesDeviceIDTelemetryPostNotFound{
				Code:    api.NewOptInt(404),
				Message: api.NewOptString("device not found or inactive"),
			}, nil
		}

		h.logger.Error("post telemetry data failed: ", err)

		return &api.DevicesDeviceIDTelemetryPostInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("internal server error"),
		}, nil
	}

	// http 200
	return &api.TelemetryData{
		DeviceId:      api.NewOptString(td.DeviceId),
		DeviceType:    api.NewOptString(td.DeviceType),
		CreatedAt:     api.NewOptDateTime(td.CreatedAt),
		TelemetryData: api.NewOptString(td.TelemetryData),
	}, nil
}
