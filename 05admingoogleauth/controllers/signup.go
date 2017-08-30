package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// SignupHandler ...
func SignupHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("SignupHandler")
		return RenderTemplate(c, "signup", nil)
	}
}
