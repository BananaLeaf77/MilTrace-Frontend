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
	StatusMoving   DeviceStatus = "moving"
	StatusStopped  DeviceStatus = "stopped"
)

type Device struct {
	ID     uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	Name   string       `gorm:"size:100;not null" json:"name"`
	IMEI   string       `gorm:"size:15;uniqueIndex" json:"imei"`
	Status DeviceStatus `gorm:"type:varchar(20);default:'inactive'" json:"status"`

	// Relationships
	Locations      []Location `gorm:"foreignKey:DeviceID" json:"-"` // Full location history
	LastLocation   *Location  `gorm:"foreignKey:ID;references:LastLocationID" json:"last_location"`
	LastLocationID *uuid.UUID `gorm:"type:uuid" json:"-"` // Pointer for optional relationship

	// Device Metrics
	CurrentSpeed   *float64 `gorm:"type:decimal(10,2)" json:"current_speed"`
	BatteryLevel   *float32 `gorm:"type:decimal(5,2)" json:"battery_level"`
	SignalStrength *int     `gorm:"type:smallint" json:"signal_strength"`

	// Timestamps
	LastHeartbeat *time.Time `json:"last_heartbeat"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type Location struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	DeviceID uuid.UUID `gorm:"type:uuid;not null;index" json:"device_id"`

	// Geospatial Data
	Latitude  float64  `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Longitude float64  `gorm:"type:decimal(9,6);not null" json:"longitude"`
	Altitude  *float64 `gorm:"type:decimal(9,2)" json:"altitude"`
	Accuracy  *float32 `gorm:"type:decimal(6,2)" json:"accuracy"`
	Heading   *float32 `gorm:"type:decimal(5,2)" json:"heading"`

	// Metadata
	Source     string    `gorm:"size:20;default:'gps'" json:"source"`
	IsMoving   bool      `gorm:"default:false" json:"is_moving"`
	RecordedAt time.Time `gorm:"index" json:"recorded_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
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
