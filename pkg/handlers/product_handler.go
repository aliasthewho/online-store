package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aliasthewho/online-store/internal/app"
	"github.com/aliasthewho/online-store/pkg/products"
	"github.com/gorilla/mux"
)

// @Summary Get all the products
// @Description Returns all the products
// @Tags products
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	productService := app.NewProductService()
	products := productService.ListProducts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// @Summary Get a product by ID
// @Description Returns a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /products/{id} [get]
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get route variables
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	productService := app.NewProductService()
	product := productService.GetProductByID(id)

	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// @Summary Create a product
// @Description Creates a product
// @Tags products
// @Accept json
// @Produce json
// @Param product body products.Product true "Product"
// @Success 201
// @Failure 400
// @Router /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product products.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	productService := app.NewProductService()
	productService.CreateProduct(product)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product created successfully")
}

// @Summary Update a product
// @Description Updates a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body products.Product true "Product"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /products/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updateProduct products.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	productService := app.NewProductService()
	err = productService.UpdateProduct(id, updateProduct)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product updated succesfully")
}

// @Summary Delete a product
// @Description Deletes a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid prouct ID", http.StatusBadRequest)
		return
	}

	productService := app.NewProductService()
	err = productService.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product delted successfully")
}
