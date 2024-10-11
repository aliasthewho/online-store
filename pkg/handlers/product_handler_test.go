package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aliasthewho/online-store/pkg/handlers"
	"github.com/gorilla/mux"
)

func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(handlers.GetProducts)

	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
