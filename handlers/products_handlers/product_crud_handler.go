package product_handlers

import (
	"api/data"
	"api/utils/http/http_router_utils"
	"api/utils/json_utils"
	"net/http"
)

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func (p *ProductHandler) getOneProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_router_utils.ParseQueryParamId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get Product by id if exist
	product, err := data.GetOneProduct(&data.Product{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json_utils.SendJsonWithHandleHttpError(w, product)
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func (p *ProductHandler) getListProductsHandler(w http.ResponseWriter) {
	productList := data.GetListProduct()

	w.WriteHeader(http.StatusOK)

	json_utils.SendJsonWithHandleHttpError(w, productList)
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func (p *ProductHandler) createOneProductHandler(w http.ResponseWriter, r *http.Request) {
	newProduct := &data.Product{}

	json_utils.ReciveJsonWithHandleHttpError(w, r.Body, newProduct)

	data.AddProduct(newProduct)

	w.WriteHeader(http.StatusCreated)

	json_utils.SendJsonWithHandleHttpError(w, newProduct)
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func (p *ProductHandler) updateOneProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_router_utils.ParseQueryParamId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProduct := &data.Product{}

	json_utils.ReciveJsonWithHandleHttpError(w, r.Body, updatedProduct)

	updatedProduct.ID = id

	err = data.UpdateProduct(updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	json_utils.SendJsonWithHandleHttpError(w, updatedProduct)
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func (p *ProductHandler) deleteOneProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := http_router_utils.ParseQueryParamId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedProduct, err := data.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	json_utils.SendJsonWithHandleHttpError(w, deletedProduct)
}
