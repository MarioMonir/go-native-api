package data

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func GetProducts() Products {
	return productList
}

func GetProductById(p *Product) (*Product, int) {
	for index, product := range productList {
		if product.ID == p.ID {
			p = product
			return p, index
		}
	}
	return nil, -1
}

func AddProduct(p *Product) {
	p.ID = len(productList) + 1
	productList = append(productList, p)
}

func UpdateProduct(p *Product) *Product {
	_, index := GetProductById(p)
	productList[index] = p
	return p
}

// Mock Database
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "One shot espresso with milk",
		Price:       3,
		SKU:         "latte123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Cortado",
		Description: "Double shot espresso with milk foam",
		Price:       2.5,
		SKU:         "corto123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
