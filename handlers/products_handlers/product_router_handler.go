package product_handlers

import (
	"log"
	"net/http"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

// var entity = "product"

// // ServeHTTP is the main entry point for the handler and satisfies http handler
// func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.logger.Println("Method: ", r.Method, " | ", r.URL.Path)
//
// 	router := http_router_utils.NewEntityRouter(r, "product")
//
// 	// Get One
// 	if router.MatchGetOneRoute() {
// 		p.getOneProductHandler(w, r)
// 		return
// 	}
//
// 	// Get List
// 	if router.MatchGetListRoute() {
// 		p.getListProductsHandler(w, r)
// 		return
// 	}
//
// 	// Create One
// 	if router.MatchCreateOneRoute() {
// 		p.createOneProductHandler(w, r)
// 		return
// 	}
//
// 	// Edit One
// 	if router.MatchUpdateOneRoute() {
// 		p.updateOneProductHandler(w, r)
// 		return
// 	}
//
// 	// Delete One
// 	if router.MatchDeleteOneRoute() {
// 		p.deleteOneProductHandler(w, r)
// 		return
// 	}
//
// 	http.Error(w, "Not Implemented Switch!", http.StatusNotImplemented)
// }

// ------------------------------------------------------

func (p *ProductHandler) RegisterProductRouter(server *http.ServeMux) {
	// Get List
	server.HandleFunc("GET /product/", p.getListProductsHandler)

	// Get One
	server.HandleFunc("GET /product/{id}", p.getOneProductHandler)

	// Create One
	server.HandleFunc("POST /product/", p.createOneProductHandler)

	// Update One
	server.HandleFunc("PUT /product/{id}", p.updateOneProductHandler)

	// Delete One
	server.HandleFunc("DELETE /product/{id}", p.deleteOneProductHandler)
}
