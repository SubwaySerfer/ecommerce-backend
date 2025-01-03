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
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключаемся к базе данных
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Инициализация зависимостей
	repo := repositories.NewRepository()
	service := services.NewService(*repo)
	handler := handlers.NewHandler(service)

	// Создаем маршруты
	r := mux.NewRouter()
	r.HandleFunc("/api/furniture", handler.GetFurnitureList).Methods("GET")
	r.HandleFunc("/api/furniture", handler.AddFurniture).Methods("POST")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	// Запускаем сервер
	log.Printf("Starting server on %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
