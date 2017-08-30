package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// LogoutHandler ...
func LogoutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LogoutHandler")

		ClearAuth(c)
		SetFlashSuccess(c, "logout successfully")
		return Redirect(c, "/")

	}
}
