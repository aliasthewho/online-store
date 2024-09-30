package v1

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter() *mux.Router {
	router := mux.NewRouter()

	RegisterUserRoutes(router)
	RegisterProductRoutes(router)

	router.Use(loggingMiddleware)
	router.Use(validateJSONMiddleware)
	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("solicitud entrante: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// validateJSONMiddleware ensures that the request has a Content-Type of application/json.
// If not, it responds with a 415 Unsupported Media Type status.
func validateJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Debe enviar JSON", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token != "mi-token-secreto" {
// 			http.Error(w, "No autorizado", http.StatusUnauthorized)
// 			return
// 		}
// 		next.ServeHTTP(w, r) // Pasar al siguiente manejador si está autenticado
// 	})
// }
