// Description: This file contains the user routes for the API.
// It registers the user routes to the router.
// It uses the handlers package to handle the requests.
// It uses the gorilla/mux package to create the router.
// It contains the following functions:
// RegisterUserRoutes: Registers the user routes to the router.
// This function registers the following routes:
// GET /users
// GET /users/{id}
// POST /users
// PUT /users/{id}
// DELETE /users/{id}
// This function uses the following handlers:
// GetUsers: Handles the GET /users route.
// GetUserByID: Handles the GET /users/{id} route.
// CreateUser: Handles the POST /users route.
// UpdateUser: Handles the PUT /users/{id} route.
// DeleteUser: Handles the DELETE /users/{id} route.
// The RegisterUserRoutes function is called in the api/v1/api.go file.
package v1

import (
	"github.com/aliasthewho/online-store/pkg/handlers"
	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers user routes
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
