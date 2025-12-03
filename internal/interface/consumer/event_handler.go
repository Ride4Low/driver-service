package consumer

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/bytedance/sonic"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ride4Low/contracts/events"
	"github.com/ride4Low/contracts/pkg/rabbitmq"
	"github.com/ride4Low/driver-service/internal/application/service"
)

type EventHandler struct {
	consumer  *rabbitmq.Consumer
	driverSvc service.DriverService
	publisher *rabbitmq.Publisher
}

func NewEventHandler(consumer *rabbitmq.Consumer, driverSvc service.DriverService, publisher *rabbitmq.Publisher) *EventHandler {
	return &EventHandler{
		consumer:  consumer,
		driverSvc: driverSvc,
		publisher: publisher,
	}
}

func (h *EventHandler) Handle(ctx context.Context, msg amqp.Delivery) error {
	var message events.AmqpMessage

	if msg.Body == nil {
		return fmt.Errorf("message body is nil")
	}

	if err := sonic.Unmarshal(msg.Body, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %v", err)
	}

	switch msg.RoutingKey {
	case events.TripEventCreated:
		return h.handleFindAndNotifyDrivers(ctx, message.Data)
	default:
		return fmt.Errorf("unknown routing key: %s", msg.RoutingKey)
	}

}

func (h *EventHandler) handleFindAndNotifyDrivers(ctx context.Context, data []byte) error {
	var payload events.TripEventData

	if err := sonic.Unmarshal(data, &payload); err != nil {
		return fmt.Errorf("failed to unmarshal message: %v", err)
	}

	suitableIDs, err := h.driverSvc.FindAvailableDrivers(ctx, payload.Trip.SelectedFare.PackageSlug)
	if err != nil {
		return fmt.Errorf("failed to find available drivers: %v", err)
	}

	if len(suitableIDs) == 0 {
		// Notify the notifier service that no drivers are available
		if err := h.publisher.PublishMessage(ctx, events.TripEventNoDriversFound, events.AmqpMessage{
			OwnerID: payload.Trip.UserID,
		}); err != nil {
			log.Printf("Failed to publish message to exchange: %v", err)
			return err
		}

		return nil
	}

	randomIndex := rand.Intn(len(suitableIDs))

	suitableDriverID := suitableIDs[randomIndex]

	// Notify the driver about a potential trip
	if err := h.publisher.PublishMessage(ctx, events.DriverCmdTripRequest, events.AmqpMessage{
		OwnerID: suitableDriverID,
		Data:    data,
	}); err != nil {
		log.Printf("Failed to publish message to exchange: %v", err)
		return err
	}

	return nil

}
