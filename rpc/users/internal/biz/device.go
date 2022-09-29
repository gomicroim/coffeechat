package biz

import (
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"users/internal/data"
	"users/internal/data/ent"
)

type DeviceUseCase struct {
	repo data.DeviceRepo
	log  *log.Logger
}

func NewDeviceUseCase(deviceRepo data.DeviceRepo, logger *log.Logger) *DeviceUseCase {
	return &DeviceUseCase{repo: deviceRepo, log: logger}
}

func (d *DeviceUseCase) Register(ctx context.Context, device *data.Device) (*data.Device, error) {
	po, err := d.repo.FindByDeviceId(ctx, device.DeviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			return d.repo.Create(ctx, device)
		}
		return nil, err
	}

	err = d.repo.UpdateByDevice(ctx, device.DeviceID, device)
	if err != nil {
		return nil, err
	}
	device.ID = po.ID
	return device, nil
}
