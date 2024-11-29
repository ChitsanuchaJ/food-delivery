package handlers

import (
	"fmt"
	"food-delivery-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type orderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) OrderHandler {
	return orderHandler{orderService}
}

func (h orderHandler) PlaceOrder(c echo.Context) error {
	fmt.Println("Place order called")

	orderReq := services.OrderRequest{}

	if err := c.Bind(&orderReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	orders, err := h.orderService.PlaceOrder(orderReq)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, orders)
}
