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

func (s *deviceServiceHandler) RegisterNewDevice(ctx context.Context, deviceData *domain.Device) error {
	err := s.repo.RegisterNewDevice(ctx, deviceData)
	if err != nil {
		return err
	}

	return nil
}

func (s *deviceServiceHandler) UpdateDevice(ctx context.Context, deviceData *domain.Device) error {
	err := s.repo.UpdateDevice(ctx, deviceData)
	if err != nil {
		return err
	}

	return nil
}

func (s *deviceServiceHandler) DeleteDevice(ctx context.Context, deviceUUID *uuid.UUID) error {
	err := s.repo.DeleteDevice(ctx, deviceUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *deviceServiceHandler) GetAllDeviceData(ctx context.Context) (*[]domain.Device, error) {
	devices, err := s.repo.GetAllDeviceData(ctx)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (s *deviceServiceHandler) GetDevice(ctx context.Context, deviceUUID *uuid.UUID) (*domain.Device, error) {
	device, err := s.repo.GetDevice(ctx, deviceUUID)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (s *deviceServiceHandler) ReceiveLocationData(ctx context.Context, deviceUUID *uuid.UUID, locationData *domain.Location) error {
	err := s.repo.ReceiveLocationData(ctx, deviceUUID, locationData)
	if err != nil {
		return err
	}

	return nil
}
