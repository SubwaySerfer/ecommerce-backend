package furniture

import (
	"database/sql"
	"fmt"
)

func CreateFurnitureTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS furniture (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2) NOT NULL,
			image_url TEXT,
			category TEXT,
			colors TEXT[],
			sizes TEXT[],
			images TEXT[]
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Furniture table created successfully!")
	return nil
}
