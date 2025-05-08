package services

import (
	"MilTrace/domain"
	"context"

	"github.com/google/uuid"
)

type deviceServiceHandler struct {
	repo domain.DeviceRepository
}

func NewDeviceService(repo domain.DeviceRepository) domain.DeviceService {
	return &deviceServiceHandler{
		repo: repo,
	}
}

func (r *deviceServiceHandler) RegisterNewDevice(deviceData *domain.Device, ctx context.Context) error {
	err := r.repo.RegisterNewDevice(deviceData, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *deviceServiceHandler) UpdateDevice(deviceData *domain.Device, ctx context.Context) error {
	err := r.repo.UpdateDevice(deviceData, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *deviceServiceHandler) DeleteDevice(deviceUUID *uuid.UUID, ctx context.Context) error {
	err := r.repo.DeleteDevice(deviceUUID, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *deviceServiceHandler) GetAllDeviceData(ctx context.Context) (*[]domain.Device, error) {
	devices, err := r.repo.GetAllDeviceData(ctx)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *deviceServiceHandler) GetDevice(deviceUUID *uuid.UUID, ctx context.Context) (*domain.Device, error) {
	device, err := r.repo.GetDevice(deviceUUID, ctx)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (r *deviceServiceHandler) ReceiveLocationData(deviceUUID *uuid.UUID, locationData *domain.Location, ctx context.Context) error {
	err := r.repo.ReceiveLocationData(deviceUUID, locationData, ctx)
	if err != nil {
		return err
	}

	return nil
}
