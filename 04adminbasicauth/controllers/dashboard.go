package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// DashboardHandler ...
func DashboardHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("DashboardHandler")
		if !IsLogin(c) {
			SetFlashError(c, "you need to login")
			return Redirect(c, "/")
		}
		return RenderTemplate(c, "dashboard", nil)
	}
}
