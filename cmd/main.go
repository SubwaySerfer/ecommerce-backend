package main

import (
	"log"
	"net/http"

	"ecommerce_backend/internal/handlers"
	"ecommerce_backend/pkg/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize HTTP server
	http.HandleFunc("/api/furniture", handlers.FurnitureHandler)
	http.HandleFunc("/api/favorites", handlers.FavoritesHandler)
	http.HandleFunc("/api/cart", handlers.CartHandler)

	log.Printf("Starting server on %s:%s", cfg.Host, cfg.Port)
	if err := http.ListenAndServe(cfg.Host+":"+cfg.Port, nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
