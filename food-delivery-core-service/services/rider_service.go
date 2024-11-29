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

	msg := fmt.Sprintf("OrderID: %v \"picked up\"", pickUpOrderReq.OrderID)
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

func (s riderService) DeliverOrder(deliverOrderReq DeliverOrderRequest) (deliverOrderResp *DeliverOrderResponse, err error) {
	fmt.Println("CORE - DeliverOrderRequest:", deliverOrderReq)

	msg := fmt.Sprintf("OrderID: %v \"delivered\"", deliverOrderReq.OrderID)
	fmt.Println(msg)
	fmt.Println("")

	event := events.OrderDelivery{
		RiderId:  deliverOrderReq.RiderID,
		OrderId:  deliverOrderReq.OrderID,
		OptField: events.OptionalField{Message: msg},
	}

	fmt.Println("Publish event:", event.GetTopicName())
	fmt.Println("")

	err = s.eventProducer.Produce(event)
	if err != nil {
		return nil, err
	}

	deliverOrderResp = &DeliverOrderResponse{
		Status: events.ORDER_STATUS_DELIVERED,
	}

	return deliverOrderResp, nil
}
