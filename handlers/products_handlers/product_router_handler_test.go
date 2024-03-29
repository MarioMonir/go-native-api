package product_handlers_test

import (
	"api/data"
	"api/utils/http/http_client_utils"
	"api/utils/http/http_server_utils"
	"api/utils/load_env_utils"
	"api/utils/logger_utils"
	"io"
	"net/http"
	"testing"
)

// ------------------------------------------------------

var (
	logger     = logger_utils.NewLogger()
	httpClient = http_client_utils.NewHttpClient(logger)
)

// ------------------------------------------------------

func init() {
	load_env_utils.LoadEnv()
	go http_server_utils.LaunchHttpServer()
}

// ------------------------------------------------------

func TestGetListProductHandler(t *testing.T) {
	products := &data.Products{}

	payload := &http_client_utils.HttpClientPayload{
		Method:       http.MethodGet,
		Endpoint:     "/product",
		EntityStruct: products,
	}

	res, err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Failed to Get List")
	}
}

// ------------------------------------------------------

func TestGetOneProductHandler(t *testing.T) {
	product := &data.Product{}

	payload := &http_client_utils.HttpClientPayload{
		Method:       http.MethodGet,
		Endpoint:     "/product/1",
		EntityStruct: product,
	}

	res, err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Failed to Get One")
	}
}

// ------------------------------------------------------

func TestCreateOneProduct(t *testing.T) {
	product := &data.Product{
		Name:        "test_name",
		Price:       0.5,
		Description: "Test Description",
		SKU:         "test_sku",
	}

	payload := &http_client_utils.HttpClientPayload{
		Method:   http.MethodPost,
		Endpoint: "/product/",
		Body:     product,
	}

	res, err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Error("Failed to Create One")
	}
}

// ------------------------------------------------------

func TestUpdateOneProduct(t *testing.T) {
	product := &data.Product{
		Name:        "test_name",
		Price:       0.5,
		Description: "Test Description",
		SKU:         "test_sku",
	}

	payload := &http_client_utils.HttpClientPayload{
		Endpoint: "/product/1",
		Body:     product,
		Method:   http.MethodPut,
	}

	res, err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Failed to Update One")
	}
}

// ------------------------------------------------------

func TestDeleteOneProduct(t *testing.T) {
	payload := &http_client_utils.HttpClientPayload{
		Endpoint: "/product/3",
		Method:   http.MethodDelete,
	}

	res, err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Failed to Update One")
	}
}
