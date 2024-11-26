package handlers

import "github.com/labstack/echo"

type NotificationHandler interface {
	SendNotification(echo.Context) error
}
