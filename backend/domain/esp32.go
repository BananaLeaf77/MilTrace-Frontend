package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Device struct {
	DeviceID   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"device_id"`
	DeviceName string     `gorm:"size:100;not null" json:"device_name"`
	Status     bool       `gorm:"default:true" json:"status"`
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
	RegisterNewDevice(deviceData *Device, ctx context.Context) error
	UpdateDevice(deviceData *Device, ctx context.Context) error
	DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(deviceUUID *uuid.UUID, ctx context.Context) (*Device, error)

	ReceiveLocationData(deviceUUID *uuid.UUID, locationData *Location, ctx context.Context) error
}

type DeviceService interface {
	RegisterNewDevice(deviceData *Device, ctx context.Context) error
	UpdateDevice(deviceData *Device, ctx context.Context) error
	DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(deviceUUID *uuid.UUID, ctx context.Context) (*Device, error)

	ReceiveLocationData(deviceUUID *uuid.UUID, locationData *Location, ctx context.Context) error
}
