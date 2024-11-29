package services

import (
	"events"
	"fmt"
	"notification-service/producer"
)

type riderService struct {
	eventProducer producer.EventProducer
}

func NewRiderService(eventProducer producer.EventProducer) RiderService {
	return riderService{eventProducer}
}

func (s riderService) PickUpOrder(pickUpOrderReq PickUpOrderRequest) (pickUpOrderResp *PickUpOrderResponse, err error) {

	fmt.Println("CORE - PickUpOrderRequest:", pickUpOrderReq)

	msg := fmt.Sprintf("OrderID: %v \"pick up\"", pickUpOrderReq.OrderID)
	fmt.Println(msg)
	fmt.Println("")

	event := events.OrderPickUp{
		RiderId:  pickUpOrderReq.RiderID,
		OrderId:  pickUpOrderReq.OrderID,
		OptField: events.OptionalField{Message: msg},
	}

	fmt.Println("Publish event:", event.GetTopicName())
	fmt.Println("")

	err = s.eventProducer.Produce(event)
	if err != nil {
		return nil, err
	}

	pickUpOrderResp = &PickUpOrderResponse{
		Status: events.ORDER_STATUS_PICKED_UP,
	}

	return pickUpOrderResp, nil
}
