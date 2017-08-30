package controllers

import (
	"log"

	"github.com/ipfans/echo-session"

	"github.com/labstack/echo"
)

// DashboardHandler ...
func DashboardHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("DashboardHandler")
		data := map[string]interface{}{}
		isLogin := false
		session := session.Default(c)
		if val := session.Get("is_login"); val != nil {
			isLogin, ok := val.(bool)
			if !ok || !isLogin {
				return c.Redirect(302, "/login")
			}
		}
		data["is_login"] = isLogin
		return c.Render(200, "dashboard", data)
	}
}
