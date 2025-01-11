package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"ecommerce_backend/internal/db/blogs"
	"ecommerce_backend/internal/db/furniture"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create furniture table if it doesn't exist
	err = furniture.CreateFurnitureTable(db)
	if err != nil {
		log.Fatal(err)
	}
	err = blogs.CreateBlogsTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
