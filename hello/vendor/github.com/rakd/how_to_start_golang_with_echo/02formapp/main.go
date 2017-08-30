package main

import (
	"html/template"
	"io"
	"log"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rakd/how_to_start_golang_with_echo/02formapp/controllers"

	"github.com/itsjamie/go-bindata-templates"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// Template is custom renderer for Echo, to render html from bindata
type Template struct {
	templates *template.Template
}

// Render renders template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplate creates a new template
func NewTemplate() *Template {
	return &Template{
		templates: binhtml.New(Asset, AssetDir).MustLoadDirectory("templates"),
	}
}

func main() {

	// Make an engine
	engine := echo.New()

	// Use precompiled embedded templates
	engine.Renderer = NewTemplate()

	// Set up echo debug level
	engine.Debug = true

	// Regular middlewares
	engine.Use(middleware.Recover())

	//csrf
	engine.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf_token",
	}))

	//session
	store := session.NewCookieStore([]byte("secret-key-kaskdfkedknjjkelaasdfjkajkeasd"))
	store.MaxAge(86400)
	engine.Use(session.Sessions("GSESSION", store))

	/*
		engine.GET("/favicon.ico", func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, "/static/images/favicon.ico")
		})
	*/
	/*
		// error handling
		engine.HTTPErrorHandler = func(err error, c echo.Context) {
			if h, ok := err.(*echo.HTTPError); ok {
				log.Print(h.Code)
				log.Print(h.Error())
				switch h.Code {
				case http.StatusBadRequest:
					// 400
					c.Render(http.StatusBadRequest, "400.html", nil)
					log.Print("400")
				case http.StatusForbidden:
					// 403
					if strings.Contains(h.Error(), "nvalid csrf token") {
						c.Render(http.StatusForbidden, "403_invalid_csrf_token.html", nil)
					} else {
						c.Render(http.StatusForbidden, "403.html", nil)
					}
				default:
					log.Print("500")
					c.Render(http.StatusInternalServerError, "500.html", nil)
				}
			}
		}
	*/

	engine.GET("/", controllers.WelcomeHandler())
	engine.GET("/login", controllers.LoginHandler())
	engine.GET("/logout", controllers.LogoutHandler())
	engine.GET("/dashboard", controllers.DashboardHandler())
	engine.POST("/login", controllers.LoginPostHandler())
	//engine.GET("/", controllers.WelcomeHandler())
	engine.Start(":3000")
	return
}
