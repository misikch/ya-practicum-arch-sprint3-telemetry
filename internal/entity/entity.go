package entity

import "time"

type TelemetryData struct {
	DeviceId      string    `bson:"deviceId" json:"deviceId"`
	DeviceType    string    `bson:"deviceType" json:"deviceType"`
	CreatedAt     time.Time `bson:"createdAt" json:"createdAt"`
	TelemetryData string    `bson:"telemetryData" json:"telemetryData"`
}

type Device struct {
	DeviceId  string    `bson:"device_id" json:"deviceId"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
