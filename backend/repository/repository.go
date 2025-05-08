package repository

import (
	"MilTrace/domain"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) domain.DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (r *deviceRepository) RegisterNewDevice(ctx context.Context, device *domain.Device) error {
	if err := r.db.WithContext(ctx).Create(device).Error; err != nil {
		return err
	}
	return nil
}

func (r *deviceRepository) UpdateDevice(ctx context.Context, device *domain.Device) error {
	if err := r.db.WithContext(ctx).Save(device).Error; err != nil {
		return err
	}
	return nil
}

func (r *deviceRepository) DeleteDevice(ctx context.Context, deviceID *uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Device{}, deviceID).Error; err != nil {
		return err
	}
	return nil
}

func (r *deviceRepository) GetAllDeviceData(ctx context.Context) (*[]domain.Device, error) {
	var devices []domain.Device
	if err := r.db.WithContext(ctx).Find(&devices).Error; err != nil {
		return nil, err
	}
	return &devices, nil
}

func (r *deviceRepository) GetDevice(ctx context.Context, deviceID *uuid.UUID) (*domain.Device, error) {
	var device domain.Device
	if err := r.db.WithContext(ctx).First(&device, "device_id = ?", deviceID).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) ReceiveLocationData(ctx context.Context, deviceID *uuid.UUID, locationData *domain.Location) error {
	if err := r.db.WithContext(ctx).Model(&domain.Device{}).Where("device_id = ?", deviceID).Association("LocationData").Append(locationData); err != nil {
		return err
	}
	return nil
}
