package blogs

import (
	"database/sql"
	"fmt"
)

func CreateBlogsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS blogs (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT,
			author VARCHAR(255) NOT NULL,
			created_at TIMESTAMP
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Blogs table created successfully!")
	return nil
}
