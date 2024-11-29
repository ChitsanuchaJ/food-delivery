package handlers

import (
	"fmt"
	"notification-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type riderHandler struct {
	riderService services.RiderService
}

func NewRiderHandler(riderService services.RiderService) RiderHandler {
	return riderHandler{riderService}
}

func (h riderHandler) PickUpOrder(c echo.Context) error {
	fmt.Println("CORE - Pick up order called")

	pickUpOrderReq := services.PickUpOrderRequest{}

	if err := c.Bind(&pickUpOrderReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	pickUpOrderResp, err := h.riderService.PickUpOrder(pickUpOrderReq)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, pickUpOrderResp)
}
