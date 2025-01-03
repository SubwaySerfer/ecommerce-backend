package services

import (
	"ecommerce_backend/internal/models"
	"ecommerce_backend/internal/repositories"
)

type Service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetFurnitureList() ([]models.Furniture, error) {
	return s.repo.GetFurnitureList()
}

// func (s *Service) AddToFavorites(userID string, furnitureID string) error {
// 	return s.repo.AddFavorite(userID, furnitureID)
// }

// func (s *Service) RemoveFromFavorites(userID string, furnitureID string) error {
// 	return s.repo.RemoveFavorite(userID, furnitureID)
// }

// func (s *Service) CreateCart(userID string, cartItems []models.CartItem) error {
// 	return s.repo.SaveCart(userID, cartItems)
// }

// func (s *Service) GetCart(userID string) ([]models.CartItem, error) {
// 	return s.repo.FetchCart(userID)
// }

func (s *Service) AddFurniture(furniture models.Furniture) error {
	return s.repo.AddFurniture(furniture)
}
