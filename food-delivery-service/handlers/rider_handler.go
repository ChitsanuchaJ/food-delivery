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

	fmt.Println("view rider called")

	riders, err := h.riderService.GetRiders()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, riders)
}
