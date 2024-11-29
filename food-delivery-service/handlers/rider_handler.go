package handlers

import (
	"fmt"
	"food-delivery-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type riderHandler struct {
	riderService services.RiderService
}

func NewRiderHandler(riderService services.RiderService) RiderHandler {
	return riderHandler{riderService}
}

func (h riderHandler) GetRiders(c echo.Context) error {
	fmt.Println("View rider called")

	riders, err := h.riderService.GetRiders()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, riders)
}

func (h riderHandler) PickUpOrder(c echo.Context) error {
	fmt.Println("Pick up order called")

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

func (h riderHandler) DeliverOrder(c echo.Context) error {
	fmt.Println("Deliver order called")

	deliverOrderReq := services.DeliverOrderRequest{}

	if err := c.Bind(&deliverOrderReq); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	deliverOrderResp, err := h.riderService.DeliverOrder(deliverOrderReq)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, deliverOrderResp)
}
