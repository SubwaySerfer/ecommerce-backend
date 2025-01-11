package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce_backend/internal/db"
	"ecommerce_backend/internal/handlers"
	"ecommerce_backend/internal/repositories"
	"ecommerce_backend/internal/services"
	"ecommerce_backend/pkg/config"

	"github.com/gorilla/mux"
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

	repo := repositories.NewRepository(database)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/furniture", handler.GetFurnitureList).Methods("GET")
	r.HandleFunc("/api/furniture", handler.AddFurniture).Methods("POST")
	r.HandleFunc("/api/blog", handler.AddBlogPost).Methods("POST")
	r.HandleFunc("/api/blog", handler.GetBlogPosts).Methods("GET")
	r.HandleFunc("/api/blog/{id}", handler.GetBlogPostByID).Methods("GET")
	r.HandleFunc("/api/blog/{id}", handler.DeleteBlogPostByID).Methods("DELETE")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	log.Printf("Starting server on %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
