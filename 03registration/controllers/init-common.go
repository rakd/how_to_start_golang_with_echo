package controllers

import (
	"fmt"
	"log"

	"regexp"

	"github.com/ipfans/echo-session"
	"github.com/rakd/how_to_start_golang_with_echo/03registration/models"

	"github.com/labstack/echo"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func init() {

}

// RenderTemplate ...
func RenderTemplate(c echo.Context, tmpl string, datatmp interface{}) error {
	log.Print("RenderTemplate")
	data := map[string]interface{}{}
	list, ok := datatmp.(map[string]interface{})
	if ok {
		for n, v := range list {
			data[n] = v
		}
	}
	data["flash_error"] = GetFlashError(c)
	data["flash_warning"] = GetFlashWarning(c)
	data["flash_info"] = GetFlashInfo(c)
	data["flash_success"] = GetFlashSuccess(c)
	data["is_login"] = IsLogin(c)

	//data["my_user_id"] = GetMyUserIDStr(c)
	//data["my_email"] = GetMyEmail(c)
	csrf := c.Get("csrf")
	csrfToken, ok := csrf.(string)
	if ok && csrfToken != "" {
		data["csrf_token"] = csrfToken
	}

	data["current_uri"] = c.Request().URL.Path

	locale := c.Get("locale")
	localeStr, ok := locale.(string)
	if ok && localeStr != "" {
		data["locale"] = localeStr
	} else {
		data["locale"] = "en"
	}

	return c.Render(200, tmpl, data)

}

const flashKeyInfo = "flash_key_info"
const flashKeyError = "flash_key_Error"
const flashKeyWarning = "flash_key_warning"
const flashKeySuccess = "flash_key_success"

func setSessionString(c echo.Context, key, value string) error {

	session := session.Default(c)
	if value == "" {
		return nil
	}
	session.Set(key, value)
	session.Save()
	return nil
}

// Redirect ...
func Redirect(c echo.Context, url string) error {

	return c.Redirect(302, url)

}

// getSessionString ...
func getSessionString(c echo.Context, key string) string {
	session := session.Default(c)
	obj := session.Get(key)
	log.Print(obj)
	if val, ok := obj.(string); ok {
		return string(val)
	}
	return ""
}

// SetFlashError ...
func SetFlashError(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyError)
}

func setFlash(c echo.Context, msg, key string) error {
	log.Print("setFlash")
	//log.Print(msg)
	//log.Print(key)
	if msg == "" {
		//log.Print("")
		return nil
	}
	session := session.Default(c)
	log.Print("")
	session.Set(key, msg)
	log.Print("")
	err := session.Save()
	if err != nil {
		log.Print(err)
	}
	log.Print("")
	return err
}

func getFlash(c echo.Context, key string) string {
	session := session.Default(c)
	obj := session.Get(key)
	if msg, ok := obj.(string); ok {
		session.Delete(key)
		session.Save()
		return msg
	}
	return ""
}

// SetFlashSuccess ...
func SetFlashSuccess(c echo.Context, msg string) error {
	log.Print("SetFlashSuccess")
	log.Print(msg)

	return setFlash(c, msg, flashKeySuccess)
}

// GetMyUserID ...
func GetMyUserID(c echo.Context) int64 {
	session := session.Default(c)
	if userID := session.Get("user_id"); userID != nil {
		val, ok := userID.(int64)
		if ok {
			return val
		}
	}
	return 0
}

// GetMyUserIDStr ...
func GetMyUserIDStr(c echo.Context) string {
	userID := GetMyUserID(c)
	if userID > 0 {
		return fmt.Sprintf("%d", userID)
	}
	return ""
}

// IsLogin ...
func IsLogin(c echo.Context) bool {

	isLogin := c.Get("is_login")
	isLoginBool, ok := isLogin.(bool)

	if ok && isLoginBool {
		log.Print("")
		return true
	}

	session := session.Default(c)
	flag := session.Get("is_login")
	if flag != nil {
		val, ok := flag.(int)
		if ok && val == 1 {
			return true
		}
	}
	return false
}

// SetAuth ...
func SetAuth(c echo.Context, user models.User) {
	log.Print("SetAuth")
	session := session.Default(c)
	session.Set("user_id", user.ID)
	session.Set("email", user.Email)
	session.Set("is_login", 1)
	session.Save()
}

// SetFlashInfo ...
func SetFlashInfo(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyInfo)
}

// GetMyEmail ...
func GetMyEmail(c echo.Context) string {
	session := session.Default(c)
	if fname := session.Get("email"); fname != nil {
		val, ok := fname.(string)
		if ok {
			return val
		}
	}
	return ""
}

// GetFlashError ...
func GetFlashError(c echo.Context) string {
	return getFlash(c, flashKeyError)
}

// GetFlashSuccess ...
func GetFlashSuccess(c echo.Context) string {
	return getFlash(c, flashKeySuccess)
}

// GetFlashInfo ...
func GetFlashInfo(c echo.Context) string {
	return getFlash(c, flashKeyInfo)
}

// SetFlashWarning ...
func SetFlashWarning(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyWarning)
}

// GetFlashWarning ...
func GetFlashWarning(c echo.Context) string {
	return getFlash(c, flashKeyWarning)
}

// ClearAuth ...
func ClearAuth(c echo.Context) {
	session := session.Default(c)
	session.Delete("is_login")
	session.Save()
}
