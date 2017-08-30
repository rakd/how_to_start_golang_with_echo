package controllers

import (
	"log"
	"github.com/ipfans/echo-session"
	
	"github.com/labstack/echo"
)

// LogoutHandler ...
func LogoutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LogoutHandler")
		
		session := session.Default(c)
		session.Delete("is_login")
		session.Save()
		
		return c.Redirect(302, "/")
	}
}
