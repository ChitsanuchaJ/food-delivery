package main

import (
	"context"
	"events"
	"fmt"
	"log"
	"net/http"
	"notification-service/consumer"
	"notification-service/handlers"
	"notification-service/producer"
	"notification-service/services"
	"notification-service/utils"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	kafkaProducer := utils.InitProducer()
	defer kafkaProducer.Close()
	eventProducer := producer.NewEventProducer(kafkaProducer)

	orderCreatedConsumer := initConsumer(events.GROUP_ORDER_CREATE)
	defer orderCreatedConsumer.Close()

	orderAcceptedConsumer := initConsumer(events.GROUP_ORDER_ACCEPT)
	defer orderAcceptedConsumer.Close()

	orderPickedUpConsumer := initConsumer(events.GROUP_ORDER_PICK_UP)
	defer orderPickedUpConsumer.Close()

	orderDeliveredConsumer := initConsumer(events.GROUP_ORDER_DELIVERY)
	defer orderDeliveredConsumer.Close()

	notiService := services.NewNotificationService()
	notiHandler := handlers.NewNotificationHandler(notiService)

	restaurantService := services.NewRestaurantService(eventProducer)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	eventHandler := consumer.NewEventHandler(notiService)
	consumerHandler := consumer.NewConsumeHandler(eventHandler)

	// Act like Order Service
	go comsumerListener(events.TOPIC_ORDER_CREATE, orderCreatedConsumer, consumerHandler)

	// Act like Rider Service
	go comsumerListener(events.TOPIC_ORDER_ACCEPT, orderAcceptedConsumer, consumerHandler)
	// go comsumerListener(events.TOPIC_ORDER_PICK_UP, orderPickedUpConsumer, consumerHandler)
	// go comsumerListener(events.TOPIC_ORDER_DELIVERY, orderDeliveredConsumer, consumerHandler)

	//////////////////////////////////////////////////////

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.POST("/notification/send", notiHandler.SendNotification)
	app.POST("/restaurant/order/accept", restaurantHandler.AcceptOrder)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := app.Start(":8001"); err != nil && err != http.ErrServerClosed {
			fmt.Println("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func comsumerListener(topic string, consumerGroup sarama.ConsumerGroup, handler sarama.ConsumerGroupHandler) {
	for {
		consumerGroup.Consume(context.Background(), []string{topic}, handler)
	}
}

func initConsumer(consumerGroupID string) sarama.ConsumerGroup {
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9093"}, consumerGroupID, nil)
	if err != nil {
		panic(err)
	}
	return consumer
}
