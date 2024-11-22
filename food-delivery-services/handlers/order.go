package handlers

import "github.com/labstack/echo"

type OrderHandler interface {
	PlaceOrder(echo.Context) error
}
