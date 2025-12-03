package config

import (
	"context"

	"github.com/ride4Low/driver-service/internal/interface/grpc/handler"
)

// Container holds all dependencies
type Container struct {
	DriverHandler *handler.DriverHandler
}

func NewContainer(ctx context.Context) (*Container, error) {

	// Initialize handlers
	driverHandler := handler.NewDriverHandler()

	return &Container{
		DriverHandler: driverHandler,
	}, nil
}
