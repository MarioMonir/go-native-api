package product_handlers

import (
	"api/utils/http/http_router_utils"
	"log"
	"net/http"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

var entity = "product"

// ServeHTTP is the main entry point for the handler and satisfies http handler
func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Method: ", r.Method, " | ", r.URL.Path)

	router := http_router_utils.NewEntityRouter(r, "product")

	// Get One
	if router.MatchGetOneRoute() {
		p.getOneProductHandler(w, r)
		return
	}

	// Get List
	if router.MatchGetListRoute() {
		p.getListProductsHandler(w)
		return
	}

	// Create One
	if router.MatchCreateOneRoute() {
		p.createOneProductHandler(w, r)
		return
	}

	// Edit One
	if router.MatchUpdateOneRoute() {
		p.updateOneProductHandler(w, r)
		return
	}

	// Delete One
	if router.MatchDeleteOneRoute() {
		p.deleteOneProductHandler(w, r)
		return
	}

	http.Error(w, "Not Implemented Switch!", http.StatusNotImplemented)
}
