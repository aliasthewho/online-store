package app

import (
	"errors"

	"github.com/aliasthewho/online-store/pkg/products"
)

type ProductService struct {
	products []products.Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []products.Product{
			{ID: 1, Name: "Laptopz", Price: 1500},
			{ID: 2, Name: "Smartphone", Price: 800},
		},
	}
}

func (ps *ProductService) ListProducts() []products.Product {
	return ps.products
}

func (ps *ProductService) GetProductByID(id int) *products.Product {
	for _, p := range ps.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func (ps *ProductService) CreateProduct(product products.Product) {
	ps.products = append(ps.products, product)
}

func (ps *ProductService) UpdateProduct(id int, updateProduct products.Product) error {
	for i, p := range ps.products {
		if p.ID == id {
			ps.products[i] = updateProduct
			return nil
		}
	}
	return errors.New("product not found")
}

func (ps *ProductService) DeleteProduct(id int) error {
	for i, p := range ps.products {
		if p.ID == id {
			ps.products = append(ps.products[:i], ps.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
