package http_client_utils

import (
	"api/utils/load_env_utils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	load_env_utils.LoadEnv()
}

type HttpClient struct {
	logger *log.Logger
}

type HttpClientPayload struct {
	Endpoint     string
	Method       string
	Body         any
	EntityStruct any
}

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func NewHttpClient(logger *log.Logger) *HttpClient {
	return &HttpClient{logger}
}

func (h *HttpClient) Query(p *HttpClientPayload) error {
	apiUrl := os.Getenv("API_URL")

	var (
		res *http.Response
		err error
		url = apiUrl + p.Endpoint
	)

	var jsonData []byte

	if p.Body != nil {
		jsonData, err = json.Marshal(p.Body)
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
		}
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
