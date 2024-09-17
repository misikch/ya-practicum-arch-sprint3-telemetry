package handler

import (
	"context"
	"device-manager/api"
)

func (h Handler) DevicesDeviceIDTelemetryGet(
	ctx context.Context,
	params api.DevicesDeviceIDTelemetryGetParams,
) (api.DevicesDeviceIDTelemetryGetRes, error) {
	res, err := h.telemetryService.GetTelemetryHistory(ctx, params.DeviceID, params.From, params.To)
	if err != nil {
		h.logger.Error("Get telemetry history failed", err)

		return &api.DevicesDeviceIDTelemetryGetInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("storage failed"),
		}, nil
	}

	if len(res) == 0 {
		return &api.DevicesDeviceIDTelemetryGetNotFound{
			Code:    api.NewOptInt(404),
			Message: api.NewOptString("empty telemetry history"),
		}, nil
	}

	out := make([]api.TelemetryData, 0, len(res))

	for _, tdata := range res {
		out = append(out, api.TelemetryData{
			DeviceId:   api.NewOptString(tdata.DeviceId),
			DeviceType: api.NewOptString(tdata.DeviceType),
			CreatedAt:  api.NewOptDateTime(tdata.CreatedAt),
			//TelemetryData: &TelemetryDataTelemetryData{},
		})
	}

	s := api.DevicesDeviceIDTelemetryGetOKApplicationJSON(out)

	return &s, nil
}

type TelemetryDataTelemetryData struct {
	Key string `json:"key"` // замените "key" на реальные поля
}
