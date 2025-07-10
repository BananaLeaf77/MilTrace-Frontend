package domain

import (
	"context"
	"time"
)

type Device struct {
	DeviceID  string     `gorm:"primaryKey;size:100" json:"device_id"`
	Latitude  float64    `gorm:"type:float" json:"latitude"`                     // Current position
	Longitude float64    `gorm:"type:float" json:"longitude"`                    // Current position
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`               // Last update time
	Locations []Location `gorm:"foreignKey:DeviceID" json:"locations,omitempty"` // History
}

type Location struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceID  string    `gorm:"size:100;index" json:"device_id"` // FK to Device
	Latitude  float64   `gorm:"type:float" json:"latitude"`
	Longitude float64   `gorm:"type:float" json:"longitude"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // Record time
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
