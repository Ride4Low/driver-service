package service

import (
	"context"

	"github.com/sithu-go/ride-share/driver-service/internal/domain/repository"
)

type DriverService interface {
	FindAvailableDrivers(ctx context.Context, packageType string) error
}

type driverService struct {
	driverRepository repository.DriverRepository
}
