package main

import (
	"log"
	"net/http"

	"ecommerce_backend/internal/db"
	"ecommerce_backend/pkg/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	cfg := config.LoadConfig()

	// Initialize HTTP server
	// http.HandleFunc("/api/furniture", handlers.FurnitureHandler)
	// http.HandleFunc("/api/favorites", handlers.FavoritesHandler)
	// http.HandleFunc("/api/cart", handlers.CartHandler)

	log.Printf("Starting server on %s:%s", cfg.Port, cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
