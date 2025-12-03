package config

import (
	"context"

	"github.com/ride4Low/driver-service/internal/application/service"
	"github.com/ride4Low/driver-service/internal/infrastructure/ephemeral/inmem"
	"github.com/ride4Low/driver-service/internal/interface/grpc/handler"
)

// Container holds all dependencies
type Container struct {
	DriverHandler *handler.DriverHandler
}

func NewContainer(ctx context.Context) (*Container, error) {

	// Initialize repositories
	driverRepo := inmem.NewDriverRepository()

	// Initialize services
	driverService := service.NewDriverService(driverRepo)

	// Initialize handlers
	driverHandler := handler.NewDriverHandler(driverService)

	return &Container{
		DriverHandler: driverHandler,
	}, nil
}
