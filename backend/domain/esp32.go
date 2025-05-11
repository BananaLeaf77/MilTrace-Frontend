package domain

import (
	"context"
	"time"
)

type Device struct {
	DeviceID  string     `gorm:"size:100;not null" json:"device_id"`
	Latitude  float64    `gorm:"type:float;not null" json:"latitude"`
	Longitude float64    `gorm:"type:float;not null" json:"longitude"`
	UpdatedAt *time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}

type DeviceRepository interface {
	RegisterNewDevice(ctx context.Context, deviceData *Device) error
	UpdateDevice(ctx context.Context, deviceData *Device) error
	DeleteDevice(ctx context.Context, deviceID string) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(ctx context.Context, deviceID string) (*Device, error)

	ReceiveLocationData(ctx context.Context, payload *Device) error
}

type DeviceService interface {
	RegisterNewDevice(ctx context.Context, deviceData *Device) error
	UpdateDevice(ctx context.Context, deviceData *Device) error
	DeleteDevice(ctx context.Context, deviceID string) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(ctx context.Context, deviceID string) (*Device, error)

	ReceiveLocationData(ctx context.Context, payload *Device) error
}
