package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/aliasthewho/online-store/api/v1"
	"github.com/aliasthewho/online-store/config"
)

// @title Online Store API
// @version 1.0
// @description This is an online store API
// @host localhost:8080
// @BasePath /api/v1
// @Schemes http
func main() {
	cfg := config.LoadConfig()

	router := v1.SetUpRouter()

	fmt.Println("Server listening on port", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, router))
}
