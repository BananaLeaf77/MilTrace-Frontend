package repository

import (
	"MilTrace/domain"
	"context"
	"errors"
	"fmt"
	"time"

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
	var deviceCount int64

	if device.DeviceID == "" {
		return errors.New("device ID cannot be empty")
	}

	err := r.db.WithContext(ctx).Model(&domain.Device{}).Where("device_id = ?", device.DeviceID).Count(&deviceCount).Error
	if err != nil {
		return fmt.Errorf("failed to check device existence: %w", err)
	}

	if deviceCount > 0 {
		return fmt.Errorf("device with ID %s already exists", device.DeviceID)
	}

	if err = r.db.WithContext(ctx).Create(device).Error; err != nil {
		return fmt.Errorf("failed to register new device: %w", err)
	}

	return nil
}

func (r *deviceRepository) UpdateDevice(ctx context.Context, device *domain.Device) error {
	if err := r.db.WithContext(ctx).Save(device).Error; err != nil {
		return err
	}
	return nil
}

func (r *deviceRepository) DeleteDevice(ctx context.Context, deviceID string) error {
	if deviceID == "" {
		return errors.New("device ID cannot be empty")
	}

	var deviceCount int64
	err := r.db.WithContext(ctx).Model(&domain.Device{}).Where("device_id = ?", deviceID).Count(&deviceCount).Error
	if err != nil {
		return fmt.Errorf("failed to check device existence: %w", err)
	}

	if deviceCount == 0 {
		return fmt.Errorf("device with ID %s does not exist", deviceID)
	}
	// Delete the device
	if err = r.db.WithContext(ctx).Where("device_id = ?", deviceID).Delete(&domain.Device{}).Error; err != nil {
		return fmt.Errorf("failed to delete device: %w", err)
	}

	return nil
}

func (r *deviceRepository) GetAllDeviceData(ctx context.Context) (*[]domain.Device, error) {
	var devices []domain.Device
	err := r.db.WithContext(ctx).
		Find(&devices).
		Error

	if err != nil {
		return nil, err
	}
	return &devices, nil
}

func (r *deviceRepository) GetDevice(ctx context.Context, deviceID string) (*domain.Device, error) {
	var device domain.Device
	if err := r.db.WithContext(ctx).Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) ReceiveLocationData(ctx context.Context, payload *domain.Device) error {
	var deviceCount int64
	tNow := time.Now()

	err := r.db.WithContext(ctx).Model(&domain.Device{}).Where("device_id = ?", payload.DeviceID).Count(&deviceCount).Error
	if err != nil {
		return fmt.Errorf("failed to check device existence: %w", err)
	}
	if deviceCount == 0 {
		return fmt.Errorf("device with ID %s does not exist", payload.DeviceID)
	}

	if payload.Latitude == 0 || payload.Longitude == 0 {
		return errors.New("latitude and longitude cannot be zero")
	}

	payload.UpdatedAt = &tNow

	// Update the device location data
	if err := r.db.WithContext(ctx).Model(&domain.Device{}).Where("device_id = ?", payload.DeviceID).Updates(payload).Error; err != nil {
		return fmt.Errorf("failed to update device location data: %w", err)
	}
	return nil
}
