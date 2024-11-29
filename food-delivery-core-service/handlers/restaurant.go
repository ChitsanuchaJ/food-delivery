package handlers

import "github.com/labstack/echo"

type RestaurantHandler interface {
	AcceptOrder(echo.Context) error
}
