package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce_backend/internal/cloudinary"
	"ecommerce_backend/internal/db"
	"ecommerce_backend/internal/handlers"
	"ecommerce_backend/internal/repositories"
	"ecommerce_backend/internal/services"
	"ecommerce_backend/pkg/config"

	gorillahandlers "github.com/gorilla/handlers"
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
	r.HandleFunc("/api/furniture/{id}", handler.DeleteFurnitureByID).Methods("DELETE")

	r.HandleFunc("/api/blog", handler.AddBlogPost).Methods("POST")
	r.HandleFunc("/api/blog", handler.GetBlogPosts).Methods("GET")
	r.HandleFunc("/api/blog/{id}", handler.GetBlogPostByID).Methods("GET")
	r.HandleFunc("/api/blog/{id}", handler.DeleteBlogPostByID).Methods("DELETE")

	r.HandleFunc("/api/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods("PUT")

	r.HandleFunc("/api/contact", handler.AddContactFormItem).Methods("POST")
	r.HandleFunc("/api/contact", handler.GetAllContactForms).Methods("GET")

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	corsOptions := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	cld, ctx := cloudinary.Credentials()
	cloudinary.UploadImage(cld, ctx)
	cloudinary.GetAssetInfo(cld, ctx)
	cloudinary.TransformImage(cld, ctx)

	log.Printf("Server is running on http://localhost:%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, corsOptions(r)); err != nil {
		log.Fatal(err)
	}
}
