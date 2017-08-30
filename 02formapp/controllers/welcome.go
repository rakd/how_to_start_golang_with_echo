package controllers

import (
	"log"

	"github.com/ipfans/echo-session"

	"github.com/labstack/echo"
)

// WelcomeHandler ...
func WelcomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("WelcomeHandler")
		data := map[string]interface{}{}
		data["is_login"] = false

		session := session.Default(c)
		if val := session.Get("is_login"); val != nil {
			isLogin, ok := val.(bool)
			if ok {
				data["is_login"] = isLogin
			}
		}

		return c.Render(200, "welcome", data)
	}
}
