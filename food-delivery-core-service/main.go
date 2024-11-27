package main

import (
	"context"
	"events"
	"notification-service/consumer"
	"notification-service/handlers"
	"notification-service/services"

	"github.com/IBM/sarama"
	"github.com/labstack/echo"
)

func main() {
	orderCreatedConsumer := initConsumer(events.GROUP_ORDER_CREATED)
	defer orderCreatedConsumer.Close()

	notiService := services.NewNotificationService()
	notiHandler := handlers.NewNotificationHandler(notiService)

	eventHandler := consumer.NewEventHandler()
	consumerHandler := consumer.NewConsumeHandler(eventHandler)

	go comsumerListener(orderCreatedConsumer, consumerHandler)

	e := echo.New()

	e.POST("/notification/send", notiHandler.SendNotification)

	e.Logger.Fatal(e.Start(":8001"))
}

func comsumerListener(consumerGroup sarama.ConsumerGroup, handler sarama.ConsumerGroupHandler) {
	topics := []string{
		events.TOPIC_ORDER_CREATED,
		events.TOPIC_ORDER_ACCEPTED,
		events.TOPIC_ORDER_PICKED_UP,
		events.TOPIC_ORDER_DELIVERED,
	}

	for {
		consumerGroup.Consume(context.Background(), topics, handler)
	}
}

func initConsumer(consumerGroupID string) sarama.ConsumerGroup {
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9093"}, consumerGroupID, nil)
	if err != nil {
		panic(err)
	}
	return consumer
}
