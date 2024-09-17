package handler

import (
	"context"
	"fmt"

	"device-manager/api"
)

func (h Handler) DevicesDeviceIDTelemetryPost(
	ctx context.Context,
	req *api.DevicesDeviceIDTelemetryPostReq,
	params api.DevicesDeviceIDTelemetryPostParams,
) (api.DevicesDeviceIDTelemetryPostRes, error) {
	fmt.Println(params)

	return nil, nil
}
