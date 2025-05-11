package services

import (
	"MilTrace/domain"
	"context"
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

func (s *deviceServiceHandler) DeleteDevice(ctx context.Context, deviceID string) error {
	err := s.repo.DeleteDevice(ctx, deviceID)
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

func (s *deviceServiceHandler) GetDevice(ctx context.Context, deviceID string) (*domain.Device, error) {
	device, err := s.repo.GetDevice(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (s *deviceServiceHandler) ReceiveLocationData(ctx context.Context, payload *domain.Device) error {
	err := s.repo.ReceiveLocationData(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
