package main

import (
	"events"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type NotificationRequest struct {
	Recipient string `json:"recipient"`
	OrderID   string `json:"order_id"`
	Message   string `json:"message"`
}

type NotificationResponse struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.POST("/notification/send", sendNotificationFromApi)

	e.Logger.Fatal(e.Start(":8001"))
}

func sendNotificationFromApi(c echo.Context) error {
	notiReq := NotificationRequest{}

	if err := c.Bind(&notiReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if notiReq.Recipient == events.RECIPIENT_CUSTOMER {
		fmt.Println("call from OrderService to customer")
	}

	notiResp := NotificationResponse{
		Status: "sent",
	}

	return c.JSON(http.StatusOK, notiResp)
}

// func sendNotificationFromKafka() error {

// }
