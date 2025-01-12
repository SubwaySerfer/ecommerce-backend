package services

import (
	"ecommerce_backend/internal/models"
	"ecommerce_backend/internal/repositories"
)

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetFurnitureList() ([]models.Furniture, error) {
	return s.repo.GetFurnitureList()
}

func (s *Service) DeleteFurnitureByID(id string) error {
	return s.repo.DeleteFurnitureByID(id)
}

func (s *Service) GetBlogPosts() ([]models.Blog, error) {
	return s.repo.GetBlogPosts()
}

func (s *Service) GetBlogPostByID(id string) (models.Blog, error) {
	return s.repo.GetBlogPostByID(id)
}

func (s *Service) DeleteBlogPostByID(id string) error {
	return s.repo.DeleteBlogPostByID(id)
}

func (s *Service) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
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

func (s *Service) AddBlogPost(blog models.Blog) error {
	return s.repo.AddBlogPost(blog)
}
