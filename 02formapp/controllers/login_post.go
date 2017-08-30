package controllers

import (
	"log"

	"github.com/ipfans/echo-session"

	"github.com/labstack/echo"
)

// LoginPostHandler ...
func LoginPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LoginPostHandler")

		session := session.Default(c)
		session.Set("is_login", true)
		session.Save()

		return c.Redirect(302, "/dashboard")
	}
}
