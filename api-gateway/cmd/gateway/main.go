package main

import (
	"log"
	"net/http"

	"api-gateway/internal/config"
	"api-gateway/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New(cfg)

	addr := ":" + cfg.Port
	log.Printf("api-gateway listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

