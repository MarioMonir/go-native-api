package product_handlers_test

import (
	"api/data"
	product_handlers "api/handlers/products_handlers"
	"api/utils/http_client_utils"
	"api/utils/http_server_utils"
	"api/utils/load_env_utils"
	"api/utils/logger_utils"
	"net/http/httptest"
	"testing"
)

var (
	logger         = logger_utils.NewLogger()
	productHandler = product_handlers.NewProductHandler(logger)
	server         = httptest.NewServer(productHandler)
	httpClient     = http_client_utils.NewHttpClient(logger)
)

func init() {
	load_env_utils.LoadEnv()
	go http_server_utils.LaunchHttpServer()
}

func TestGetListProductHandler(t *testing.T) {
	products := &data.Products{}

	payload := &http_client_utils.HttpClientPayload{
		Endpoint:     "/product",
		EntityStruct: products,
	}

	err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
}

func TestGetOneProductHandler(t *testing.T) {
	product := &data.Product{}

	payload := &http_client_utils.HttpClientPayload{
		Endpoint:     "/product/1",
		EntityStruct: product,
	}

	err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateOneProduct(t *testing.T) {
	httpClient = http_client_utils.NewHttpClient(logger)

	product := &data.Product{
		Name:        "test_name",
		Price:       0.5,
		Description: "Test Description",
		SKU:         "test_sku",
	}

	payload := &http_client_utils.HttpClientPayload{
		Endpoint: "/product/1",
		Body:     product,
	}

	err := httpClient.Query(payload)
	if err != nil {
		t.Error(err)
	}
}
