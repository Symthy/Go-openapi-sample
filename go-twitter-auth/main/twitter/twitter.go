package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/Symth/golang-practices/go-twitter-auth/main/yaml"
	"github.com/garyburd/go-oauth/oauth"
)

// Account アカウント
type Account struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
}

const (
	callbackURL = "http://localhost:8080/twitter/callback"
)

// GetConnect 接続を取得する
func buildTwitterClient() (*oauth.Client, error) {
	conf, err := yaml.LoadConfigForYaml()
	if err != nil {
		return nil, err
	}
	client := &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  conf.Twitter.ConsumerApiKey,
			Secret: conf.Twitter.ConsumerApiSecretKey,
		},
	}
	return client, nil
}

func callGetRequestTokenApi(oc *oauth.Client) (*oauth.Credentials, error) {
	reqToken, err := oc.RequestTemporaryCredentials(nil, callbackURL, nil)
	log.Print("success get request token: ", reqToken)
	return reqToken, err
}

func callGetAuthorizationUriApi(oc *oauth.Client, reqToken *oauth.Credentials) string {
	return oc.AuthorizationURL(reqToken, nil)
}

func callGetAccessTokenApi(oc *oauth.Client, reqToken *oauth.Credentials,
	oauthVerifier string) (*oauth.Credentials, error) {
	accToken, _, err := oc.RequestToken(nil, reqToken, oauthVerifier)
	log.Print("success get access token: ", accToken)
	return accToken, err
}

type TwitterAccessHandler struct {
	client oauth.Client
}

func NewTwitterAccessHandler() *TwitterAccessHandler {
	client, err := buildTwitterClient()
	if err != nil {
		log.Fatal("load Yaml os.Open err:", err)
	}
	return &TwitterAccessHandler{
		client: *client,
	}
}

func (h *TwitterAccessHandler) GetRequestToken() (*oauth.Credentials, error) {
	return callGetRequestTokenApi(&h.client)
}

func (h *TwitterAccessHandler) GetAuthorizationUrl(reqToken *oauth.Credentials) (*string, error) {
	oc := &h.client
	url := callGetAuthorizationUriApi(oc, reqToken)
	return &url, nil
}

func (h *TwitterAccessHandler) GetAccessToken(reqToken *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := &h.client
	return callGetAccessTokenApi(oc, reqToken, oauthVerifier)
}

func (h *TwitterAccessHandler) GetUserAllInfo(accessToken *oauth.Credentials) (*interface{}, error) {
	var user interface{}
	err := h.callGetApi(
		accessToken,
		"https://api.twitter.com/1.1/account/verify_credentials.json",
		url.Values{"include_email": {"true"}},
		&user)
	if err != nil {
		return nil, fmt.Errorf("failure get account info: %s", err.Error())
	}
	return &user, nil
}

func (h *TwitterAccessHandler) GetAccount(accessToken *oauth.Credentials) (*Account, error) {
	var user Account
	err := h.callGetApi(
		accessToken,
		"https://api.twitter.com/1.1/account/verify_credentials.json",
		url.Values{"include_email": {"true"}},
		&user)
	if err != nil {
		return nil, fmt.Errorf("failure get account info: %s", err.Error())
	}
	return &user, nil
}

func (h *TwitterAccessHandler) callGetApi(cred *oauth.Credentials, urlStr string, form url.Values,
	data interface{}) error {
	oc := &h.client
	resp, err := oc.Get(nil, cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != http.StatusOK {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("url: %s, status: %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
