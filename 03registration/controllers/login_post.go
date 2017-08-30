package controllers

import (
	"log"

	"github.com/labstack/echo"
	"github.com/rakd/how_to_start_golang_with_echo/03registration/models"
)

// LoginPostHandler ...
func LoginPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("LoginPostHandler")
		var err error
		email := c.Request().FormValue("email")
		password := c.Request().FormValue("password")

		user := models.User{}
		user.Email = email
		user.Password = password
		log.Print("")
		user, err = user.Login()
		log.Print("")
		if err != nil {
			log.Print("")
			SetFlashError(c, err.Error())
			return Redirect(c, "/login")
		}
		log.Print("")

		SetAuth(c, user)
		SetFlashSuccess(c, "login ok")
		return Redirect(c, "/dashboard")
	}
}
