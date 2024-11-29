package handlers

import "github.com/labstack/echo"

type RiderHandler interface {
	GetRiders(echo.Context) error
	PickUpOrder(echo.Context) error
	DeliverOrder(echo.Context) error
}
