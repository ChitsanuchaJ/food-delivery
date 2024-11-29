package handlers

import (
	"fmt"
	"food-delivery-service/services"

	"net/http"

	"github.com/labstack/echo"
)

type menuHandler struct {
	menuService services.MenuService
}

func NewMenuHandler(menuService services.MenuService) MenuHandler {
	return menuHandler{menuService}
}

func (h menuHandler) GetMenusByID(c echo.Context) error {
	fmt.Println("View menu called")

	restaurantId := c.Param("id")

	menus, err := h.menuService.GetMenu(restaurantId)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, menus)
}
