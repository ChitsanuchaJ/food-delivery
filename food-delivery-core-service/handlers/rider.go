package handlers

import "github.com/labstack/echo"

type RiderHandler interface {
	PickUpOrder(echo.Context) error
	DeliverOrder(echo.Context) error
}
