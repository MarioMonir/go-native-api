package product_handlers

import (
	"api/data"
	"api/utils/json_utils"
	"net/http"
	"strconv"
)

func (u *ProductHandler) getProductByIdHandler(w http.ResponseWriter, r *http.Request) {
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
	json_utils.SendJson(w, product)
}

func (u *ProductHandler) getProductsHandler(w http.ResponseWriter, _ *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Encode json", http.StatusInternalServerError)
		return
	}
}

func (u *ProductHandler) addProductHandler(w http.ResponseWriter, r *http.Request) {
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Decode json", http.StatusInternalServerError)
	}

	data.AddProduct(newProduct)

	json_utils.SendJson(w, newProduct)
}

func (u *ProductHandler) updateProductHandler(w http.ResponseWriter, r *http.Request) {
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
	err = newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Decode json", http.StatusInternalServerError)
	}

	newProduct.ID = id
	data.UpdateProduct(newProduct)
	json_utils.SendJson(w, newProduct)
}
