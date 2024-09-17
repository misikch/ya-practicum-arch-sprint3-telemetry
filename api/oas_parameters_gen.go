// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DevicesDeviceIDTelemetryGetParams is parameters of GET /devices/{device_id}/telemetry operation.
type DevicesDeviceIDTelemetryGetParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDTelemetryGetParams(packed middleware.Parameters) (params DevicesDeviceIDTelemetryGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDTelemetryGetParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDTelemetryGetParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DevicesDeviceIDTelemetryLatestGetParams is parameters of GET /devices/{device_id}/telemetry/latest operation.
type DevicesDeviceIDTelemetryLatestGetParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDTelemetryLatestGetParams(packed middleware.Parameters) (params DevicesDeviceIDTelemetryLatestGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDTelemetryLatestGetParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDTelemetryLatestGetParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DevicesDeviceIDTelemetryPostParams is parameters of POST /devices/{device_id}/telemetry operation.
type DevicesDeviceIDTelemetryPostParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDTelemetryPostParams(packed middleware.Parameters) (params DevicesDeviceIDTelemetryPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDTelemetryPostParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDTelemetryPostParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
