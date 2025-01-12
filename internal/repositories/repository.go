package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

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
	query := `SELECT id, name, description, price, image_url, category, colors, sizes, images FROM furniture`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query furniture: %w", err)
	}
	defer rows.Close()

	var furnitureList []models.Furniture
	for rows.Next() {
		var furniture models.Furniture
		var colors, sizes, images pq.StringArray
		if err := rows.Scan(&furniture.ID, &furniture.Name, &furniture.Description, &furniture.Price, &furniture.ImageURL, &furniture.Category, &colors, &sizes, &images); err != nil {
			return nil, fmt.Errorf("failed to scan furniture: %w", err)
		}
		furniture.Colors = colors
		furniture.Sizes = sizes
		furniture.Images = images
		furnitureList = append(furnitureList, furniture)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if len(furnitureList) == 0 {
		return nil, errors.New("no furniture available")
	}

	return furnitureList, nil
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

func (r *Repository) DeleteFurnitureByID(id string) error {
	query := `DELETE FROM furniture WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete furniture: %w", err)
	}
	return nil
}

func (r *Repository) AddBlogPost(blog models.Blog) error {
	query := `
			INSERT INTO blogs
			(title, content, author, created_at)
			VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		query,
		blog.Title,
		blog.Content,
		blog.Author,
		blog.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert blog post: %w", err)
	}
	return nil
}

func (r *Repository) GetBlogPosts() ([]models.Blog, error) {
	query := `SELECT id, title, content, author, created_at FROM blogs`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query blog posts: %w", err)
	}
	defer rows.Close()

	var blogPosts []models.Blog
	for rows.Next() {
		var blog models.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan blog post: %w", err)
		}
		blogPosts = append(blogPosts, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if len(blogPosts) == 0 {
		return nil, errors.New("no blog posts available")
	}

	return blogPosts, nil
}

func (r *Repository) GetBlogPostByID(id string) (models.Blog, error) {
	query := `SELECT id, title, content, author, created_at FROM blogs WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var blog models.Blog
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt); err != nil {
		return models.Blog{}, fmt.Errorf("failed to scan blog post: %w", err)
	}

	return blog, nil
}

func (r *Repository) DeleteBlogPostByID(id string) error {
	// Check if the blog post exists
	existsQuery := `SELECT 1 FROM blogs WHERE id = $1`
	var exists int
	err := r.db.QueryRow(existsQuery, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("blog post not found")
		}
		return fmt.Errorf("failed to check blog post existence: %w", err)
	}

	// Delete the blog post
	deleteQuery := `DELETE FROM blogs WHERE id = $1`
	_, err = r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete blog post: %w", err)
	}
	return nil
}

func (r *Repository) CreateUser(user models.User) error {
	query := `
		INSERT INTO users
		(username, email, password, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Password,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

// func (r *Repository) UpdateUser(user models.User) error {
// 	query := `
// 		UPDATE users
// 		SET username = $2, email = $3, password = $4
// 		WHERE id = $1
// 	`

// 	fmt.Printf("user: %v\n", user)
// 	_, err := r.db.Exec(
// 		query,
// 		user.ID,
// 		user.Username,
// 		user.Email,
// 		user.Password,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to update user: %w", err)
// 	}
// 	return nil
// }

func (r *Repository) UpdateUser(user models.User) error {
	query := `UPDATE users SET `
	params := []interface{}{}
	paramID := 1

	if user.Username != "" {
		query += fmt.Sprintf("username = $%d, ", paramID)
		params = append(params, user.Username)
		paramID++
	}
	if user.Email != "" {
		query += fmt.Sprintf("email = $%d, ", paramID)
		params = append(params, user.Email)
		paramID++
	}
	if user.Password != "" {
		query += fmt.Sprintf("password = $%d, ", paramID)
		params = append(params, user.Password)
		paramID++
	}

	// Remove the trailing comma and space
	query = query[:len(query)-2]

	query += fmt.Sprintf(" WHERE id = $%d", paramID)
	params = append(params, user.ID)

	_, err := r.db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
