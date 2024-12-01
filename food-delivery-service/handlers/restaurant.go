package handlers

import "github.com/labstack/echo"

type RestaurantHandler interface {
	GetRestaurants(echo.Context) error
	AcceptOrder(echo.Context) error
}
