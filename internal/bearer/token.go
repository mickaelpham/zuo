package bearer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mickael/zuo/internal/conf"
)

const (
	APP_PATH   string = "/.zuo/"
	TOKEN_FILE string = APP_PATH + "access-token.json"
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
	if fromFile, valid := load(); valid {
		return fromFile
	}

	return fetch()
}

func load() (*AccessToken, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Print(err)
		return nil, false
	}

	dat, err := ioutil.ReadFile(homeDir + TOKEN_FILE)
	if err != nil {
		log.Println("no stored access token found " +
			"(I will fetch one for ya, don't worry)")
		return nil, false
	}

	var token AccessToken
	err = json.Unmarshal(dat, &token)
	if err != nil {
		log.Print(err)
		return nil, false
	}

	if time.Now().After(token.Expiry) {
		log.Print("found stored token, but it's expired " +
			"(no big deal, I will fetch a new one!)")
		return nil, false
	}

	return &token, true
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

	message, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(homeDir+TOKEN_FILE, message, 0644)
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
