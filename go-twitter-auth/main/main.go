package main

import (
	"log"
	"net/http"

	"github.com/Symth/golang-practices/go-twitter-auth/main/ses"
	"github.com/Symth/golang-practices/go-twitter-auth/main/twitter"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// e.GET("/login", login(e))
	// e.GET("/logout", logout())
	// e.POST("/check", check())

	rh := newRestHandler()

	e.GET("/twitter/token", rh.getRequestToken())
	e.GET("/twitter/auth", rh.authorize())
	e.GET("/twitter/callback", rh.getAccessToken())
	// /twitter/auth 実行後に実行要
	e.GET("/twitter/account", rh.getAccount())
	e.GET("/twitter/account/all", rh.getUserAllInfo())
	e.Logger.Fatal(e.Start(":8080"))
}

type RestHandler struct {
	twitterAccessor *twitter.TwitterAccessHandler
	sessionAccessor *ses.SessionAccessor
}

func newRestHandler() *RestHandler {
	return &RestHandler{
		twitterAccessor: twitter.NewTwitterAccessHandler(),
		sessionAccessor: ses.NewSessionAccessor(),
	}
}

func (h *RestHandler) getRequestToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken, err := h.twitterAccessor.GetRequestToken()
		if err != nil {
			c.Error(err)
		}
		return c.JSON(http.StatusOK, reqToken)
	}
}

func (h *RestHandler) authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken, err := h.twitterAccessor.GetRequestToken()
		if err != nil {
			c.Error(err)
		}
		url, err := h.twitterAccessor.GetAuthorizationUrl(reqToken)
		if err != nil {
			c.Error(err)
		}
		h.sessionAccessor.SetRequestTokenToSession(c, reqToken)
		return c.Redirect(http.StatusFound, *url)
	}
}

// callback
func (h *RestHandler) getAccessToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("twitter authorization success")

		oauthToken := c.QueryParam("oauth_token")
		log.Print("oauth_token: ", oauthToken)
		oauthVerifier := c.QueryParam("oauth_verifier")
		log.Print("oauth_verify: ", oauthVerifier)

		reqToken, err := h.sessionAccessor.GetRequestTokenInSession(c)
		log.Print("request token in session: ", reqToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		accToken, err := h.twitterAccessor.GetAccessToken(reqToken, oauthVerifier)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		h.sessionAccessor.SetAccessTokenToSession(c, accToken)
		return c.JSON(http.StatusOK, accToken)
	}
}

func (h *RestHandler) getAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		accToken, err := h.sessionAccessor.GetAccessTokenInSession(c)
		log.Print("access token in session: ", accToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		account, err := h.twitterAccessor.GetAccount(accToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, account)
	}
}

func (h *RestHandler) getUserAllInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		accToken, err := h.sessionAccessor.GetAccessTokenInSession(c)
		log.Print("access token in session: ", accToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		account, err := h.twitterAccessor.GetUserAllInfo(accToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, account)
	}
}
