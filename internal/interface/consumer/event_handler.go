package consumer

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ride4Low/contracts/events"
	"github.com/ride4Low/contracts/pkg/rabbitmq"
)

type EventHandler struct {
	consumer *rabbitmq.Consumer
}

func NewEventHandler(consumer *rabbitmq.Consumer) *EventHandler {
	return &EventHandler{
		consumer: consumer,
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

	return nil

}
