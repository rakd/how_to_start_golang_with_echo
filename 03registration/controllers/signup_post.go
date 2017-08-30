package controllers

import (
	"log"

	"github.com/rakd/how_to_start_golang_with_echo/03registration/models"

	"github.com/labstack/echo"
)

// SignupPostHandler ...
func SignupPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("SignupPostHandler")

		email := c.Request().FormValue("email")
		password := c.Request().FormValue("password")
		password2 := c.Request().FormValue("password2")
		if password != password2 {
			SetFlashError(c, "wrong password")
			return Redirect(c, "/signup")
		}
		log.Print("")
		user, err := models.GetUserByEmail(email)
		log.Print("")
		if err == nil && user.ID > 0 {
			log.Print("")
			SetFlashError(c, "the email is alraedy used. ")
			return Redirect(c, "/signup")
		}
		log.Print("")

		user = models.NewUser()
		user.Email = email
		user.Password = password
		log.Print("")
		user, err = user.Create()
		log.Print("")
		if err != nil {
			log.Print("")
			SetFlashError(c, err.Error())
			log.Print("")
			return Redirect(c, "/signup")
		}
		log.Print("")

		SetFlashSuccess(c, "signup successfully, you can login")
		return Redirect(c, "/login")
	}
}
