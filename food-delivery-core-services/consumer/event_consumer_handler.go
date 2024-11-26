package consumer

import (
	"encoding/json"
	"events"
	"fmt"
)

type EventHandler interface {
	Handle(string, []byte)
}

type eventHandler struct{}

func NewEventHandler() EventHandler {
	return eventHandler{}
}

func (e eventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case events.TOPIC_ORDER_CREATED:
		event := &events.OrderCreate{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("@@@ eventBytes =>", eventBytes)
		fmt.Println("@@@ event =>", event)
		// Do something
	case events.TOPIC_ORDER_ACCEPTED:
		event := &events.OrderAccept{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something
	case events.TOPIC_ORDER_PICKED_UP:
		event := &events.OrderPickUp{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something
	case events.TOPIC_ORDER_DELIVERED:
		event := &events.OrderDelivery{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something
	}

}
