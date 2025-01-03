package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"ecommerce_backend/internal/models"

	"github.com/lib/pq"
)

type Repository struct {
	furnitureList []models.Furniture
	favoriteList  []models.Furniture
	cartList      []models.CartItem
	mu            sync.RWMutex
	db            *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:            db,
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

//	func (r *Repository) AddFurniture(furniture models.Furniture) error {
//		r.mu.Lock()
//		defer r.mu.Unlock()
//		r.furnitureList = append(r.furnitureList, furniture)
//		return nil
//	}
func (r *Repository) AddFurniture(furniture models.Furniture) error {
	query := `
		INSERT INTO furniture
		(name, description, price, image_url, category, colors, sizes, images)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(
		query,
		furniture.Name,
		furniture.Description,
		furniture.Price,
		furniture.ImageURL,
		furniture.Category,
		pq.Array(furniture.Colors), // Используем pq.Array для массивов
		pq.Array(furniture.Sizes),
		pq.Array(furniture.Images),
	)
	if err != nil {
		return fmt.Errorf("failed to insert furniture: %w", err)
	}
	return nil
}
