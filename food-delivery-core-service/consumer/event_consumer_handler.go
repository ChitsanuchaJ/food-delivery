package consumer

import (
	"encoding/json"
	"events"
	"fmt"
	"notification-service/services"
)

type EventHandler interface {
	Handle(string, []byte)
}

type eventHandler struct {
	notiService services.NotificationService
}

func NewEventHandler(notiService services.NotificationService) EventHandler {
	return eventHandler{notiService}
}

func (e eventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case events.TOPIC_ORDER_CREATE:
		event := &events.OrderCreate{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Consume topic", events.TOPIC_ORDER_CREATE, "log =>", event)
		fmt.Println("")

		notiReq := services.NotificationRequest{
			Recipient: events.RECIPIENT_RESTAURANT,
			OrderID:   event.OrderId,
			Message:   event.OptField.Message,
		}
		e.notiService.SendNotification(notiReq)

	case events.TOPIC_ORDER_ACCEPT:
		event := &events.OrderAccept{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Consume topic", events.TOPIC_ORDER_ACCEPT, "log =>", event)
		fmt.Println("")

		notiReq := services.NotificationRequest{
			Recipient: events.RECIPIENT_RIDER,
			OrderID:   event.OrderId,
			Message:   event.OptField.Message,
		}
		e.notiService.SendNotification(notiReq)

	case events.TOPIC_ORDER_PICK_UP:
		event := &events.OrderPickUp{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something

	case events.TOPIC_ORDER_DELIVERY:
		event := &events.OrderDelivery{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something

	default:
		fmt.Print("Invalid topic.")
	}
}
