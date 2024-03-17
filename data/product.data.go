package data

import (
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	filteredProductList := []*Product{}
	for _, p := range productList {
		if p.DeletedAt == "" {
			filteredProductList = append(filteredProductList, p)
		}
	}
	return filteredProductList
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

func UpdateProduct(p *Product) {
	_, index := GetProductById(p)
	p.UpdatedAt = time.Now().String()
	productList[index] = p
}

func DeleteProduct(p *Product) {
	_, index := GetProductById(p)
	p.UpdatedAt = time.Now().String()
	p.DeletedAt = time.Now().String()
	productList[index] = p
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
		Name:        "Espresso",
		Description: "Espresso Coffe",
		Price:       2,
		SKU:         "espresso123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
	{
		ID:          3,
		Name:        "Cortado",
		Description: "Double shot espresso with milk foam",
		Price:       2.5,
		SKU:         "corto123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
