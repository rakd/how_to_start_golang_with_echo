package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// LoginHandler ...
func LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LoginHandler")

		return RenderTemplate(c, "login", nil)
	}
}
