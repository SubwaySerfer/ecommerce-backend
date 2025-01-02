package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"ecommerce_backend/pkg/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
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
