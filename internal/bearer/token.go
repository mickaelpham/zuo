package bearer

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/mickael/zuo/internal/conf"
)

type AccessToken struct {
	Val    string
	Expiry time.Time
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Jti         string `json:"jti"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func Token() *AccessToken {
	conf := conf.FromEnv()

	startReq := time.Now().Unix()
	resp, err := http.PostForm(
		conf.BaseUrl+"/oauth/token",
		url.Values{
			"client_id":     {conf.ClientId},
			"client_secret": {conf.ClientSecret},
			"grant_type":    {"client_credentials"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var body tokenResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	return &AccessToken{
		Val:    body.AccessToken,
		Expiry: time.Unix(startReq+body.ExpiresIn, 0),
	}
}
