package handlers

import (
	"ecommerce_backend/internal/models"
	"ecommerce_backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetFurnitureList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetFurnitureList\n")
	furnitureList, err := h.Service.GetFurnitureList()
	if err != nil {
		fmt.Printf("Error getting furniture list: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(furnitureList); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func (h *Handler) ToggleFavorite(w http.ResponseWriter, r *http.Request) {
// 	var payload struct {
// 		ID string `json:"id"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err := h.Service.ToggleFavorite(payload.ID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }

// func (h *Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
// 	var payload models.Cart
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err := h.Service.CreateCart(payload)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

func (h *Handler) AddFurniture(w http.ResponseWriter, r *http.Request) {
	var furniture models.Furniture
	if err := json.NewDecoder(r.Body).Decode(&furniture); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.AddFurniture(furniture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) AddBlogPost(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.AddBlogPost(blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
