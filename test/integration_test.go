package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/aliasthewho/online-store/api/v1"
)

func TestGetProducts(t *testing.T) {
	router := v1.SetUpRouter()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status OK, but got %v", response.Code)
	}
}
