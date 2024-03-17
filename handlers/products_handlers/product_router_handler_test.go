package product_handlers_test

import (
	"api/data"
	product_handlers "api/handlers/products_handlers"
	"api/utils/http_client_utils"
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

func TestGetListProductHandler(t *testing.T) {
	products := &data.Products{}
	httpClient.Get(t, products, "/product")
}

func TestGetOneProductHandler(t *testing.T) {
	product := &data.Product{}
	httpClient.Get(t, product, "/product/1")
}
