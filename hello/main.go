package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rakd/how_to_start_golang_with_echo/hello/controllers"
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

	engine.GET("/", controllers.WelcomeHandler())
	engine.Start(":3000")
	return
}
