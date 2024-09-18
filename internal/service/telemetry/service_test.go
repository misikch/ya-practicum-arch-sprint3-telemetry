package telemetry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"device-manager/internal/entity"
)

func TestService_AddTelemetry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockStorage(ctrl)
	mockDatabus := NewMockDatabusProducer(ctrl)
	service := NewService(mockStorage, mockDatabus)

	ctx := context.TODO()
	data := entity.TelemetryData{
		DeviceId:      "device123",
		DeviceType:    "sensor",
		CreatedAt:     time.Now(),
		TelemetryData: "telemetry_info",
	}

	activeDevice := &entity.Device{
		DeviceId: "device123",
		Status:   "active",
	}

	type testCase struct {
		name        string
		mock        func()
		expectedErr error
	}

	tests := []testCase{
		{
			name: "Success",
			mock: func() {
				mockStorage.EXPECT().GetDeviceById(ctx, data.DeviceId).Return(activeDevice, nil)
				mockStorage.EXPECT().AddTelemetry(ctx, data.DeviceId, data.DeviceType, data.CreatedAt, data.TelemetryData).Return(nil)
				mockDatabus.EXPECT().PublishTelemetry(ctx, data).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "DeviceNotActive",
			mock: func() {
				inactiveDevice := &entity.Device{
					DeviceId: "device123",
					Status:   "inactive",
				}
				mockStorage.EXPECT().GetDeviceById(ctx, data.DeviceId).Return(inactiveDevice, nil)
			},
			expectedErr: ErrDeviceNotActive,
		},
		{
			name: "GetDeviceByIdError",
			mock: func() {
				mockStorage.EXPECT().GetDeviceById(ctx, data.DeviceId).Return(nil, errors.New("database error"))
			},
			expectedErr: errors.New("failed to get device: database error"),
		},
		{
			name: "AddTelemetryError",
			mock: func() {
				mockStorage.EXPECT().GetDeviceById(ctx, data.DeviceId).Return(activeDevice, nil)
				mockStorage.EXPECT().AddTelemetry(ctx, data.DeviceId, data.DeviceType, data.CreatedAt, data.TelemetryData).Return(errors.New("database error"))
			},
			expectedErr: errors.New("failed to add telemetry: database error"),
		},
		{
			name: "PublishTelemetryError",
			mock: func() {
				mockStorage.EXPECT().GetDeviceById(ctx, data.DeviceId).Return(activeDevice, nil)
				mockStorage.EXPECT().AddTelemetry(ctx, data.DeviceId, data.DeviceType, data.CreatedAt, data.TelemetryData).Return(nil)
				mockDatabus.EXPECT().PublishTelemetry(ctx, data).Return(errors.New("databus error"))
			},
			expectedErr: errors.New("failed to publish telemetry in databus: databus error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := service.AddTelemetry(ctx, data)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
