package handlers

import "github.com/labstack/echo"

type MenuHandler interface {
	GetMenusByID(echo.Context) error
}
