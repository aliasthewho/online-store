// Description: This file contains the product routes for the API.
// It registers the product routes to the router.
// It uses the handlers package to handle the requests.
// It uses the gorilla/mux package to create the router.
// It contains the following functions:
// RegisterProductRoutes: Registers the product routes to the router.
// This function registers the following routes:
// GET /products
// GET /products/{id}
// POST /products
// PUT /products/{id}
// DELETE /products/{id}
// This function uses the following handlers:
// GetProducts: Handles the GET /products route.
// GetProductByID: Handles the GET /products/{id} route.
// CreateProduct: Handles the POST /products route.
// UpdateProduct: Handles the PUT /products/{id} route.
// DeleteProduct: Handles the DELETE /products/{id} route.
// The RegisterProductRoutes function is called in the api/v1/api.go file.
package v1

import (
	"github.com/aliasthewho/online-store/pkg/handlers"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers product routes
func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
}
