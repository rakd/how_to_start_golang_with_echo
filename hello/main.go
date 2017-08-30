package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	// Make an engine
	engine := echo.New()

	// Set up echo debug level
	engine.Debug = true

	// Regular middlewares
	engine.Use(middleware.Recover())

	engine.GET("/", WelcomeHandler())
	engine.Start(":3000")
	return

}

// WelcomeHandler ...
func WelcomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.HTML(200, "adfadasda")
	}
}
