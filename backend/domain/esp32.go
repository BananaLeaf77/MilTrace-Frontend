package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Device struct {
	DeviceID   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"device_id"`
	DeviceName string     `gorm:"size:100;not null" json:"device_name"`
	Status     bool       `gorm:"default:false" json:"status"`
	Location   *Location  `gorm:"foreignKey:DeviceID" json:"location,omitempty"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type Location struct {
	LocationID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"location_id"`
	DeviceID   uuid.UUID `gorm:"type:uuid;not null" json:"device_id"`
	Latitude   float64   `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Longitude  float64   `gorm:"type:decimal(9,6);not null" json:"longitude"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type DeviceRepository interface {
	RegisterNewDevice(ctx context.Context, deviceData *Device) error
	UpdateDevice(ctx context.Context, deviceData *Device) error
	DeleteDevice(ctx context.Context, deviceUUID *uuid.UUID) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(ctx context.Context, deviceUUID *uuid.UUID) (*Device, error)

	ReceiveLocationData(ctx context.Context, deviceUUID *uuid.UUID, locationData *Location) error
}

type DeviceService interface {
	RegisterNewDevice(ctx context.Context, deviceData *Device) error
	UpdateDevice(ctx context.Context, deviceData *Device) error
	DeleteDevice(ctx context.Context, deviceUUID *uuid.UUID) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(ctx context.Context, deviceUUID *uuid.UUID) (*Device, error)

	ReceiveLocationData(ctx context.Context, deviceUUID *uuid.UUID, locationData *Location) error
}
