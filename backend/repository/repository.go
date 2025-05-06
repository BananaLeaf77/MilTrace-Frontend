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

func (r *deviceRepository) RegisterNewDevice(deviceData *domain.Device, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(deviceData).Error
}

func (r *deviceRepository) UpdateDevice(deviceData *domain.Device, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(deviceData).Error
}

func (r *deviceRepository) DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error {
	return r.db.WithContext(ctx).Where("id = ?", deviceUUID).Delete(&domain.Device{}).Error
}

func (r *deviceRepository) GetAllDeviceData(ctx context.Context) (*[]domain.Device, error) {
	var devices []domain.Device
	err := r.db.WithContext(ctx).Find(&devices).Error
	return &devices, err
}

func (r *deviceRepository) GetDevice(deviceUUID *uuid.UUID) (*domain.Device, error) {
	var device domain.Device
	err := r.db.Where("id = ?", deviceUUID).First(&device).Error
	return &device, err
}
