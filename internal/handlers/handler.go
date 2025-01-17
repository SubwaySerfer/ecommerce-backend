package handlers

import (
	"ecommerce_backend/internal/models"
	"ecommerce_backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

func handleServiceError(w http.ResponseWriter, err error, message string) {
	fmt.Printf("%s: %v\n", message, err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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

func (h *Handler) DeleteFurnitureByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeleteFurnitureByID\n")
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.Service.DeleteFurnitureByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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

func (h *Handler) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetBlogPosts\n")
	blogPosts, err := h.Service.GetBlogPosts()

	if err != nil {
		handleServiceError(w, err, "Error getting blog posts")
		return
	}
	writeJSONResponse(w, http.StatusOK, blogPosts)
}

func (h *Handler) GetBlogPostByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetBlogPostByID\n")
	vars := mux.Vars(r)
	id := vars["id"]

	blogPost, err := h.Service.GetBlogPostByID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "Blog post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve blog post", http.StatusInternalServerError)
		}
		return
	}
	writeJSONResponse(w, http.StatusOK, blogPost)
}

func (h *Handler) DeleteBlogPostByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeleteBlogPost\n")
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeleteBlogPostByID(id)
	if err != nil {
		handleServiceError(w, err, "Error deleting blog post")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateUser\n")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.CreateUser(user)
	if err != nil {
		handleServiceError(w, err, "Error creating user")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("UpdateUser\n")
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id
	err := h.Service.UpdateUser(user)
	if err != nil {
		handleServiceError(w, err, "Error updating user")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AddContactFormItem(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AddContactForm\n")
	var contactForm models.ContactForm
	if err := json.NewDecoder(r.Body).Decode(&contactForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.AddContactFormItem(contactForm)
	if err != nil {
		handleServiceError(w, err, "Error adding contact form")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAllContactForms(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetAllContactForms\n")
	contactForms, err := h.Service.GetAllContactForms()
	if err != nil {
		handleServiceError(w, err, "Error getting contact forms")
		return
	}
	writeJSONResponse(w, http.StatusOK, contactForms)
}
