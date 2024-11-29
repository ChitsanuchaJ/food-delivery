package services

import (
	"events"
	"fmt"
	"notification-service/producer"
)

type restaurantService struct {
	eventProducer producer.EventProducer
}

func NewRestaurantService(eventProducer producer.EventProducer) RestaurantService {
	return restaurantService{eventProducer}
}

func (s restaurantService) AcceptOrder(acceptOrderReq AcceptOrderRequest) (acceptOrderResp *AcceptOrderResponse, err error) {

	msg := fmt.Sprintf("OrderID: %v \"accept\"", acceptOrderReq.OrderID)
	fmt.Println(msg)
	fmt.Println("Restaurant name:", acceptOrderReq.RestaurantName)
	fmt.Println("")

	notiMessage := fmt.Sprintf("Pick order id: %v at restaurant name: %v", acceptOrderReq.OrderID, acceptOrderReq.RestaurantName)

	event := events.OrderAccept{
		OrderId:  acceptOrderReq.OrderID,
		OptField: events.OptionalField{Message: notiMessage},
	}

	fmt.Println("Publish event:", event.GetTopicName())
	fmt.Println("")

	err = s.eventProducer.Produce(event)
	if err != nil {
		return nil, err
	}

	acceptOrderResp = &AcceptOrderResponse{
		Status: events.ORDER_STATUS_ACCEPTED,
	}

	return acceptOrderResp, nil
}
