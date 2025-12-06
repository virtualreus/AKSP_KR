package main

import (
	"context"
	"log"
	"net/http"

	"medical-service/internal/config"
	"medical-service/internal/database"
	"medical-service/internal/router"
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
	log.Printf("medical-service listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
