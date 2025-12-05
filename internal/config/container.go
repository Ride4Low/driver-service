package config

import (
	"context"

	"github.com/ride4Low/contracts/env"
	"github.com/ride4Low/contracts/pkg/rabbitmq"
	"github.com/ride4Low/driver-service/internal/application/service"
	"github.com/ride4Low/driver-service/internal/infrastructure/ephemeral/inmem"
	"github.com/ride4Low/driver-service/internal/interface/consumer"
	"github.com/ride4Low/driver-service/internal/interface/grpc/handler"
)

// Container holds all dependencies
type Container struct {
	// Infrastructure
	rmq *rabbitmq.RabbitMQ

	// Handlers
	DriverHandler *handler.DriverHandler

	// messaging
	Consumer *rabbitmq.Consumer
}

func NewContainer(ctx context.Context) (*Container, error) {
	// Initialize repositories
	driverRepo := inmem.NewDriverRepository()

	// Initialize services
	driverService := service.NewDriverService(driverRepo)

	// Initialize handlers
	driverHandler := handler.NewDriverHandler(driverService)

	// Initialize messaging
	rmq, err := rabbitmq.NewRabbitMQ(env.GetString("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/"))
	if err != nil {
		return nil, err
	}

	publisher := rabbitmq.NewPublisher(rmq)

	eventHandler := consumer.NewEventHandler(driverService, publisher)

	consumer := rabbitmq.NewConsumer(rmq, eventHandler)

	return &Container{
		DriverHandler: driverHandler,
		Consumer:      consumer,
	}, nil
}

func (c *Container) Close() error {
	return c.rmq.Close()
}
