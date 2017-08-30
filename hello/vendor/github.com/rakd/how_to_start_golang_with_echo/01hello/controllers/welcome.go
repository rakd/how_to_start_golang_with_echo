package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// WelcomeHandler ...
func WelcomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("WelcomeHandler")
		return c.HTML(200, "hello")
	}
}
