package contact_form

import (
	"database/sql"
	"fmt"
)

func CreateContactFormTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS contact_form (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			subject VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			created TIMESTAMP NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Contact Form table created successfully!")
	return nil
}
