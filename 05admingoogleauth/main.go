package main

import (
	"html/template"
	"io"
	"log"

	session "github.com/ipfans/echo-session"
	"github.com/itsjamie/go-bindata-templates"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rakd/how_to_start_golang_with_echo/05admingoogleauth/controllers"
	"github.com/rakd/how_to_start_golang_with_echo/05admingoogleauth/models"
)

func init() {
	log.SetFlags(log.Lshortfile)
	models.Setup()
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
	engine.GET("/signup", controllers.SignupHandler())
	engine.POST("/signup", controllers.SignupPostHandler())
	engine.GET("/logout", controllers.LogoutHandler())
	engine.GET("/dashboard", controllers.DashboardHandler())
	engine.POST("/login", controllers.LoginPostHandler())

	adminGroup := engine.Group("/admin")
	//basic auth for admin page
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "test" && password == "test" {
			return true, nil
		}
		return false, nil
	}))

	// google auth for admin
	adminGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			adminEmails := map[string]bool{}
			adminEmails["rakd0930@gmail.com"] = true

			session := session.Default(c)
			googleEmailTmp := session.Get("google_oauth_email")
			googleEmail, ok := googleEmailTmp.(string)
			if !ok || googleEmailTmp == nil || googleEmail == "" {
				log.Print("need to login via google oauth")
				return c.Redirect(302, "/auth/google/login")
			}
			if !adminEmails[googleEmail] {
				//NG, non admin google email
				return c.Redirect(302, "/auth/google/login")
			}
			// OK
			//before

			err := next(c)
			//after

			return err
		}
	})

	adminGroup.GET("", controllers.AdminIndexHandler())
	adminGroup.GET("/", controllers.AdminIndexHandler())

	// cool custom logging
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
	}))

	engine.GET("/auth/google/login", controllers.AuthGoogleLoginHandler())
	engine.GET("/auth/google/callback", controllers.AuthGoogleCallbackHandler())
	engine.GET("/auth/google/logout", controllers.AuthGoogleLogoutHandler())

	engine.Start(":3000")
	return
}
