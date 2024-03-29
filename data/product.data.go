package data

import (
	"errors"
	"time"
)

// ------------------------------------------------------

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

// ------------------------------------------------------

func GetListProduct() Products {
	filteredProductList := []*Product{}
	for _, p := range productList {
		if p.DeletedAt == "" {
			filteredProductList = append(filteredProductList, p)
		}
	}
	return filteredProductList
}

// ------------------------------------------------------

func GetProductById(id int) (*Product, int) {
	for index, product := range productList {
		if product.ID == id {
			return product, index
		}
	}
	return nil, -1
}

// ------------------------------------------------------

func GetOneProduct(p *Product) (*Product, error) {
	product, _ := GetProductById(p.ID)
	if product == nil {
		return nil, errors.New("Product not found")
	}
	return product, nil
}

// ------------------------------------------------------

func AddProduct(p *Product) {
	p.ID = len(productList) + 1
	p.CreatedAt = time.Now().UTC().String()
	p.UpdatedAt = time.Now().UTC().String()
	productList = append(productList, p)
}

// ------------------------------------------------------

func UpdateProduct(p *Product) error {
	product, index := GetProductById(p.ID)
	if product == nil {
		return errors.New("Product not found")
	}

	p.UpdatedAt = time.Now().String()
	productList[index] = p

	return nil
}

// ------------------------------------------------------

func DeleteProduct(id int) (*Product, error) {
	product, index := GetProductById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	product.UpdatedAt = time.Now().UTC().String()
	product.DeletedAt = time.Now().UTC().String()
	productList[index] = product

	return product, nil
}

// ------------------------------------------------------

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
