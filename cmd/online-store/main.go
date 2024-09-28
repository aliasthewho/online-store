package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/aliasthewho/online-store/api/v1"
	"github.com/aliasthewho/online-store/config"
)

func main() {
	cfg := config.LoadConfig()

	router := v1.SetUpRouter()

	fmt.Println("Server listening on port", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, router))
}
