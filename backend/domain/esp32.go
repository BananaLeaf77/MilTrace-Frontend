package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DeviceStatus string

const (
	StatusActive   DeviceStatus = "active"
	StatusInactive DeviceStatus = "inactive"
	StatusFaulty   DeviceStatus = "faulty"
)

type Device struct {
	ID           uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string       `gorm:"size:100;not null" json:"name"`
	IMEI         string       `gorm:"size:15;uniqueIndex" json:"imei"`
	Status       DeviceStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	LastSeenAt   *time.Time   `json:"last_seen_at"`
	BatteryLevel *float32     `json:"battery_level"`
	CreatedAt    time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

type DeviceRepository interface {
	RegisterNewDevice(deviceData *Device, ctx context.Context) error
	UpdateDevice(deviceData *Device, ctx context.Context) error
	DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(deviceUUID *uuid.UUID) (*Device, error)
}

type DeviceService interface {
	RegisterNewDevice(deviceData *Device, ctx context.Context) error
	UpdateDevice(deviceData *Device, ctx context.Context) error
	DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error
	GetAllDeviceData(ctx context.Context) (*[]Device, error)
	GetDevice(deviceUUID *uuid.UUID) (*Device, error)
}
