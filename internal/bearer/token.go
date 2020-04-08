package bearer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/mickael/zuo/internal/conf"
)

const (
	APP_PATH    string = "/.zuo/"
	TOKENS_FILE string = APP_PATH + "access-tokens.json"
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
	if fromEnv, valid := load(); valid {
		return fromEnv
	}

	return fetch()
}

func load() (*AccessToken, bool) {
	value, found := syscall.Getenv("ZUO_ACCESS_TOKEN_VALUE")
	expiryValue, envExpiry := syscall.Getenv("ZUO_ACCESS_TOKEN_EXPIRES_AT")
	expiresAt, err := strconv.ParseInt(expiryValue, 10, 64)
	expiry := time.Unix(expiresAt, 0)

	if !found || !envExpiry || err != nil || time.Now().After(expiry) {
		return nil, false
	}

	return &AccessToken{Val: value, Expiry: expiry}, true
}

func (t *AccessToken) store() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := homeDir + APP_PATH
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	message := []byte(t.Val)
	err = ioutil.WriteFile(homeDir+TOKENS_FILE, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fetch() *AccessToken {
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

	token := &AccessToken{
		Val:    body.AccessToken,
		Expiry: time.Unix(startReq+body.ExpiresIn, 0),
	}

	token.store()
	return token
}
