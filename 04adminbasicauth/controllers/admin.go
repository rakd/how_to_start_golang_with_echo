package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// AdminIndexHandler ...
func AdminIndexHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("AdminIndexHandler")
		return RenderTemplate(c, "admin_index", nil)

	}
}
