package handlers

import (
	"net/http"
	"notification-service/services"

	"github.com/labstack/echo"
)

type notificationHandler struct {
	notiService services.NotificationService
}

func NewNotificationHandler(notiService services.NotificationService) NotificationHandler {
	return notificationHandler{notiService}
}

func (h notificationHandler) SendNotification(c echo.Context) error {
	notiReq := services.NotificationRequest{}

	if err := c.Bind(&notiReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	notiResp, err := h.notiService.SendNotification(&notiReq)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, notiResp)
}
