package product_handlers

import (
	"api/utils/query_params_utils"
	"log"
	"net/http"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

var GetProductByIdRegex = query_params_utils.RegexpToGetEntityById("product")

// ServeHTTP is the main entry point for the handler and satisfies http handler
func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get One
	if r.Method == http.MethodGet && GetProductByIdRegex.MatchString(r.URL.Path) {
		p.getOneProductHandler(w, r)
		return
	}

	// Get List
	if r.Method == http.MethodGet {
		p.getListProductsHandler(w, r)
		return
	}

	// Create One
	if r.Method == http.MethodPost {
		p.createOneProductHandler(w, r)
		return
	}

	// Edit One
	if r.Method == http.MethodPut {
		p.updateOneProductHandler(w, r)
		return
	}

	// Edit One
	if r.Method == http.MethodDelete {
		p.deleteOneProductHandler(w, r)
		return
	}

	http.Error(w, "Not Implemented Switch!", http.StatusNotImplemented)
}
