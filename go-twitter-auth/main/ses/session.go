package ses

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	// session name
	sessionName      = "twitter_session"
	requestTokenKey  = "requestToken"
	requestSecretKey = "requestTokenSecret"
	accessTokenKey   = "accessToken"
	accessSecretKey  = "accessTokenSecret"
)

type SessionAccessor struct {
}

func NewSessionAccessor() *SessionAccessor {
	return &SessionAccessor{}
}

func (sa *SessionAccessor) getSession(c echo.Context) *sessions.Session {
	session, err := session.Get(sessionName, c)
	if err != nil {
		log.Fatalf("Error session, %s", err.Error())
		return nil
	}
	log.Printf("session is new: %t", session.IsNew)
	return session
}

func (sa SessionAccessor) GetRequestTokenInSession(c echo.Context) (*oauth.Credentials, error) {
	session := sa.getSession(c)
	return sa.getTokenInSession(session, requestTokenKey, requestSecretKey)
}

func (sa SessionAccessor) GetAccessTokenInSession(c echo.Context) (*oauth.Credentials, error) {
	session := sa.getSession(c)
	return sa.getTokenInSession(session, accessTokenKey, accessSecretKey)
}

func (sa *SessionAccessor) getTokenInSession(session *sessions.Session, tokenKeyName string,
	secretKeyName string) (*oauth.Credentials, error) {

	token, ok := session.Values[tokenKeyName]
	if !ok {
		log.Printf("failure get value in session: %s", tokenKeyName)
		return nil, fmt.Errorf("error get token: %s", tokenKeyName)
	}

	secret, ok := session.Values[secretKeyName]
	if !ok {
		log.Printf("failure get value in session: %s", secretKeyName)
		return nil, fmt.Errorf("error get secret:  %s", secretKeyName)
	}

	credentials := &oauth.Credentials{
		Token:  token.(string),
		Secret: secret.(string),
	}

	return credentials, nil
}

func (sa *SessionAccessor) SetRequestTokenToSession(c echo.Context, token *oauth.Credentials) {
	session := sa.getSession(c)
	sa.setValuesToSession(session, map[string]interface{}{
		requestTokenKey:  token.Token,
		requestSecretKey: token.Secret,
	})
	if err := sa.saveSession(c, session); err != nil {
		log.Print("Error saving request token in session")
	}
}

func (sa *SessionAccessor) SetAccessTokenToSession(c echo.Context, token *oauth.Credentials) {
	session := sa.getSession(c)
	sa.setValuesToSession(session, map[string]interface{}{
		accessTokenKey:  token.Token,
		accessSecretKey: token.Secret,
	})
	if err := sa.saveSession(c, session); err != nil {
		log.Print("Error saving access token in session")
	}
}

func (sa SessionAccessor) setValuesToSession(session *sessions.Session, data map[string]interface{}) error {
	for key, value := range data {
		session.Values[key] = value
	}
	return nil
}

func (sa SessionAccessor) saveSession(c echo.Context, session *sessions.Session) error {
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}

func login(e *echo.Echo) echo.HandlerFunc {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	return func(c echo.Context) error {
		sess, _ := session.Get(sessionName, c)
		sess.Options = &sessions.Options{
			//Path:でsessionの有効な範囲を指定｡指定無しで全て有効になる｡
			//有効な時間
			MaxAge: 86400 * 7,
			//trueでjsからのアクセス拒否
			HttpOnly: true,
		}
		sess.Values["test"] = "twitter_auth_test"
		sess.Values["auth"] = true
		//状態保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}

func logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		//ログアウト
		sess.Values["auth"] = false
		//状態を保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}

func check() echo.HandlerFunc {
	return func(c echo.Context) error {
		//sessionを見る
		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error")
		}
		//ログインしているか
		if b, _ := sess.Values["auth"]; b != true {
			return c.String(http.StatusUnauthorized, "401")
		} else {
			return c.String(http.StatusOK, sess.Values["test"].(string))
		}
	}
}
