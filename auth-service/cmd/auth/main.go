package main

import (
	"context"
	"log"
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer db.Close()

	if err := database.Migrate(context.Background(), db); err != nil {
		log.Fatalf("db migrate error: %v", err)
	}

	r := router.New(cfg, db)

	addr := ":" + cfg.Port
	log.Printf("auth-service listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
