package config

import (
	"context"
	"fmt"

	"github.com/ride4Low/contracts/env"
	"github.com/ride4Low/contracts/pkg/otel"
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

	// otel
	otelProvider *otel.Provider
}

func NewContainer(ctx context.Context) (*Container, error) {
	// Initialize OpenTelemetry
	otelCfg := otel.DefaultConfig("driver-service")
	otelCfg.JaegerEndpoint = env.GetString("JAEGER_ENDPOINT", "jaeger:4317")

	otelProvider, err := otel.Setup(ctx, otelCfg)
	if err != nil {
		return nil, err
	}

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
		otelProvider:  otelProvider,
	}, nil
}

func (c *Container) Close() error {
	var errs []error
	if c.otelProvider != nil {
		if err := c.otelProvider.Shutdown(context.Background()); err != nil {
			errs = append(errs, err)
		}
	}
	if c.rmq != nil {
		if err := c.rmq.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}

	return nil
}
