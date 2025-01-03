package repositories

import (
	"errors"
	"sync"

	"ecommerce_backend/internal/models"
)

type Repository struct {
	furnitureList []models.Furniture
	favoriteList  []models.Furniture
	cartList      []models.CartItem
	mu            sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		furnitureList: []models.Furniture{},
		favoriteList:  []models.Furniture{},
		cartList:      []models.CartItem{},
	}
}

func (r *Repository) GetFurnitureList() ([]models.Furniture, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.furnitureList) == 0 {
		return nil, errors.New("no furniture available")
	}
	return r.furnitureList, nil
}

func (r *Repository) AddToFavorites(furniture models.Furniture) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.favoriteList = append(r.favoriteList, furniture)
	return nil
}

func (r *Repository) RemoveFromFavorites(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, item := range r.favoriteList {
		if item.ID == id {
			r.favoriteList = append(r.favoriteList[:i], r.favoriteList[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found in favorites")
}

func (r *Repository) CreateCart(cartItem models.CartItem) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cartList = append(r.cartList, cartItem)
}

func (r *Repository) AddFurniture(furniture models.Furniture) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.furnitureList = append(r.furnitureList, furniture)
	return nil
}
