package handlers

import (
	"fmt"
	"food-delivery-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type restaurantHandler struct {
	restaurantService services.RestaurantService
}

func NewRestaurantHandler(restaurantService services.RestaurantService) RestaurantHandler {
	return restaurantHandler{restaurantService}
}

func (h restaurantHandler) GetRestaurants(c echo.Context) error {

	fmt.Println("view restaurant called")

	restaurants, err := h.restaurantService.GetRestaurants()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, restaurants)
}
