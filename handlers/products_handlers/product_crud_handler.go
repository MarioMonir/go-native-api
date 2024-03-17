package product_handlers

import (
	"api/data"
	"api/utils/json_utils"
	"net/http"
	"strconv"
)

func (p *ProductHandler) getOneProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query params using Regexp
	matches := GetProductByIdRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	// Parse the id from the query params that matched
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	// Get Product by id if exist
	product, _ := data.GetProductById(&data.Product{ID: id})
	if product == nil {
		http.Error(w, "Product Not Found!", http.StatusNotFound)
		return
	}

	json_utils.SendJsonWithHandleHttpError(w, product)
}

func (p *ProductHandler) getListProductsHandler(w http.ResponseWriter, _ *http.Request) {
	productList := data.GetProducts()
	json_utils.SendJsonWithHandleHttpError(w, productList)
}

func (p *ProductHandler) createOneProductHandler(w http.ResponseWriter, r *http.Request) {
	newProduct := &data.Product{}
	json_utils.ReciveJsonWithHandleHttpError(w, r.Body, newProduct)

	data.AddProduct(newProduct)
	json_utils.SendJsonWithHandleHttpError(w, newProduct)
}

func (p *ProductHandler) updateOneProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query params using Regexp
	matches := GetProductByIdRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	// Parse the id from the query params that matched
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	newProduct := &data.Product{}

	json_utils.ReciveJsonWithHandleHttpError(w, r.Body, newProduct)

	newProduct.ID = id
	data.UpdateProduct(newProduct)
	json_utils.SendJsonWithHandleHttpError(w, newProduct)
}

func (p *ProductHandler) deleteOneProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query params using Regexp
	matches := GetProductByIdRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	// Parse the id from the query params that matched
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	// Get Product by id if exist
	product, _ := data.GetProductById(&data.Product{ID: id})
	if product == nil {
		http.Error(w, "Product Not Found!", http.StatusNotFound)
		return
	}

	product.ID = id
	data.DeleteProduct(product)

	json_utils.SendJsonWithHandleHttpError(w, product)
}
