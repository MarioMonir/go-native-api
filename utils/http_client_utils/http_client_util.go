package http_client_utils

import (
	"api/utils/json_utils"
	"api/utils/load_env_utils"
	"log"
	"net/http"
	"os"
	"testing"
)

func init() {
	load_env_utils.LoadEnv()
}

type HttpClient struct {
	logger *log.Logger
}

func NewHttpClient(logger *log.Logger) *HttpClient {
	return &HttpClient{logger}
}

func (h *HttpClient) Get(t *testing.T, T any, endpoint string) {
	apiUrl := os.Getenv("API_URL")

	res, err := http.Get(apiUrl + endpoint)
	if err != nil {
		t.Error("error in request the api \n", err)
	}
	defer res.Body.Close()

	err = json_utils.FromJSON(res.Body, T)
	if err != nil {
		t.Error("Unable to Jsonify Res Body to the struct \n", err)
	}
}
