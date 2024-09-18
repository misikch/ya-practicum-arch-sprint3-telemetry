package handler

import (
	"context"
	"device-manager/api"
)

func (h *Handler) DevicesDeviceIDTelemetryLatestGet(
	ctx context.Context,
	params api.DevicesDeviceIDTelemetryLatestGetParams,
) (api.DevicesDeviceIDTelemetryLatestGetRes, error) {
	res, err := h.telemetryService.GetTelemetryLatest(ctx, params.DeviceID)
	if err != nil {
		h.logger.Error("Get telemetry latest failed: ", err)

		return &api.DevicesDeviceIDTelemetryLatestGetInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("storage failed"),
		}, nil
	}

	if res == nil {
		return &api.DevicesDeviceIDTelemetryLatestGetNotFound{
			Code:    api.NewOptInt(404),
			Message: api.NewOptString("empty telemetry history"),
		}, nil
	}

	// http 200
	return &api.TelemetryData{
		DeviceId:      api.NewOptString(res.DeviceId),
		DeviceType:    api.NewOptString(res.DeviceType),
		CreatedAt:     api.NewOptDateTime(res.CreatedAt),
		TelemetryData: api.NewOptString(res.TelemetryData),
	}, nil
}
