package handlers

import (
	"fmt"
	"notification-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type restaurantHandler struct {
	restaurantService services.RestaurantService
}

func NewRestaurantHandler(restaurantService services.RestaurantService) RestaurantHandler {
	return restaurantHandler{restaurantService}
}

func (h restaurantHandler) AcceptOrder(c echo.Context) error {
	fmt.Println("CORE - Accept order called")

	acceptOrderReq := services.AcceptOrderRequest{}

	if err := c.Bind(&acceptOrderReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	acceptOrderResp, err := h.restaurantService.AcceptOrder(acceptOrderReq)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, acceptOrderResp)
}
