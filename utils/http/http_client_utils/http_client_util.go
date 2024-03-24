package http_client_utils

import (
	"api/utils/load_env_utils"
	"bytes"
	"encoding/json"
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

func NewHttpClient(logger *log.Logger) *HttpClient {
	return &HttpClient{logger}
}

func (h *HttpClient) Query(p *HttpClientPayload) (*http.Response, error) {
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
			return nil, err
		}
	}

	req, err := http.NewRequest(p.Method, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
