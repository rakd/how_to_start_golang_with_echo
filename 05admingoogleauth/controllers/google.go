package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

var googleOAuthConf oauth2.Config

const googleOAuthCookieName = "google-websta-cokies"

const callbackURI = "/auth/google/callback"

func init() {
	// this is for ADMIN FOR WEBSTA, registered in ohno@a-fis.com / https://console.cloud.google.com/apis/credentials?project=websta-1320&authuser=1
	// TODO: it's better for us to replace this clientID/Secret before releasing.
	clientID := "856008728885-h4c16ld4l3lo39bupdl2off5humgsn32.apps.googleusercontent.com"
	clientSecret := "IlE01_lFmDaIwSellp-wfRPV"

	if clientID == "" || clientSecret == "" {
		googleOAuthOK = false
	}

	googleOAuthConf = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://www.googleapis.com/oauth2/v4/token",
		},
	}
	googleOAuthOK = true

}

var googleOAuthOK bool

// AuthGoogleLoginHandler ...
func AuthGoogleLoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("AuthGoogleLoginHandler")

		state := uuid.NewV4().String()
		log.Print("state=" + state)
		session := session.Default(c)
		session.Set(googleOAuthCookieName, state)
		session.Set("google_oauth_email", nil)
		uri := c.Request().URL.Path
		if strings.HasPrefix(uri, "/") && !strings.HasPrefix(uri, "/api") && !strings.HasPrefix(uri, "/auth") {
			session.Set("return_uri", uri)
		}

		session.Save()

		redierctURL := getCallbackURL(c)
		log.Print(redierctURL)
		googleOAuthConf.RedirectURL = redierctURL
		url := googleOAuthConf.AuthCodeURL(state)

		return Redirect(c, url)
	}
}

// AuthGoogleCallbackHandler ...
func AuthGoogleCallbackHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("AuthGoogleCallbackHandler")

		if googleOAuthOK == false {

			return c.HTML(200, "no setup clientID or clientSecret for google auth")
		}

		session := session.Default(c)
		session.Get(googleOAuthCookieName)
		stateOrig := session.Get(googleOAuthCookieName)

		log.Print(stateOrig)
		if stateOrig == nil || stateOrig.(string) == "" {
			//c.String(200, "no session")
			//c.Abort()
			return c.HTML(200, "no session")
		}
		state := stateOrig.(string)

		state2 := c.Request().URL.Query().Get("state")
		if state != state2 {
			//c.String(200, "state is not match. state="+state+", state2="+state2)
			//c.Abort()
			return c.HTML(200, "state is not match. state="+state+", state2="+state2)
		}

		redierctURL := getCallbackURL(c)
		log.Print(redierctURL)
		googleOAuthConf.RedirectURL = redierctURL
		// get auth code
		code := c.Request().URL.Query().Get("code")
		log.Print(code)

		//tok, err := googleOAuthConf.Exchange(c, code)
		tok, err := googleOAuthConf.Exchange(oauth2.NoContext, code)

		if err != nil {
			return c.HTML(200, "Exchange error. cannot get token from code : "+err.Error())
		}

		// check whether the token is valid
		if tok.Valid() == false {
			//c.String(200, "token is invalid.")
			//c.Abort()
			//return
			return c.HTML(200, "token is invalid.")
		}

		// get oauth2 clinet service
		// // if you don't need to get user information, we can skip it.
		service, err := v2.New(googleOAuthConf.Client(oauth2.NoContext, tok))
		if err != nil {
			return c.HTML(200, "get oauth2 client service error : "+err.Error())
		}

		// get token info
		// it has email & user id, etc. if you don't need to get it, we can skip it.
		//tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(c).Do()
		tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(oauth2.NoContext).Do()
		if err != nil {
			return c.HTML(200, "get token info error: "+err.Error())
		}
		log.Print(tokenInfo)
		log.Print(tokenInfo.Email)

		//helpers.StoreGoogleAuthEmail(c, tokenInfo.Email)
		session.Set("google_oauth_email", tokenInfo.Email)
		returnURI := session.Get("return_uri")
		session.Set("return_uri", "")
		session.Save()

		if returnURI != nil {
			if uri, ok := returnURI.(string); ok {
				if uri != "" && strings.HasPrefix(uri, "/") {
					if !strings.HasPrefix(uri, "/auth") {
						return Redirect(c, uri)
					}
				}
			}
		}

		return Redirect(c, "/admin")
	}
}

// getCallbackURL ...
func getCallbackURL(c echo.Context) string {
	log.Print("GoogleCallback")

	if strings.Contains(strings.ToLower(c.Request().Proto), "https") {
		scheme := "https"
		return fmt.Sprintf(scheme+"://%s%s", c.Request().Host, callbackURI)
	}
	scheme := "http"
	return fmt.Sprintf(scheme+"://%s%s", c.Request().Host, callbackURI)

}

// AuthGoogleLogoutHandler ...
func AuthGoogleLogoutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("GoogleLogout")

		session := session.Default(c)
		session.Clear()
		session.Set("google_oauth_email", nil)
		session.Clear()
		session.Save()

		return Redirect(c, "/")

	}
}
