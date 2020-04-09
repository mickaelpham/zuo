package command

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mickael/zuo/internal/bearer"
	"github.com/mickael/zuo/internal/conf"
)

type QueryResponse struct {
	Done         bool                `json:"done"`
	QueryLocator string              `json:"queryLocator"`
	Records      []map[string]string `json:"records"`
	Size         int                 `json:"size"`
}

type queryRequest struct {
	QueryString string `json:"queryString"`
}

func Query(zoql string) *QueryResponse {
	conf := conf.FromEnv()

	payload, err := json.Marshal(queryRequest{QueryString: zoql})
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		conf.BaseUrl+"/v1/action/query",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Bearer "+bearer.Token().Val)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var queryResponse QueryResponse
	err = json.NewDecoder(resp.Body).Decode(&queryResponse)
	if err != nil {
		log.Fatal(err)
	}

	return &queryResponse
}
