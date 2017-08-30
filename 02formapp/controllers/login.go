package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// LoginHandler ...
func LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LoginHandler")
		data := map[string]interface{}{}
		data["test"] = "testtest"
		data["test2"] = "testtest2222"

		csrf := c.Get("csrf")
		csrfToken, ok := csrf.(string)
		if ok && csrfToken != "" {
			//log.Print("")
			data["csrf_token"] = csrfToken
		}

		return c.Render(200, "login", data)
	}
}
