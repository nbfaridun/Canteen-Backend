package migrations

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS user_role (
			user_role_id SERIAL PRIMARY KEY,
			name VARCHAR(50) UNIQUE NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "user" (
			user_id SERIAL PRIMARY KEY,
			user_role_id INT NOT NULL REFERENCES user_role(user_role_id) ON DELETE CASCADE,
			username VARCHAR(50) UNIQUE NOT NULL,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			is_active BOOLEAN DEFAULT TRUE
		);`,
		`CREATE TABLE IF NOT EXISTS client_category (
			client_category_id SERIAL PRIMARY KEY,
			name VARCHAR(50) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			is_active BOOLEAN DEFAULT TRUE
		);`,
		`CREATE TABLE IF NOT EXISTS client (
			client_id SERIAL PRIMARY KEY,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			age INT NOT NULL,
			gender VARCHAR(10) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			client_category_id INT NOT NULL REFERENCES client_category(client_category_id) ON DELETE CASCADE,
			balance FLOAT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			is_active BOOLEAN DEFAULT TRUE
		);`,
		`CREATE TABLE IF NOT EXISTS session (
			session_id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
			refresh_token TEXT UNIQUE NOT NULL,
			expires_at TIMESTAMP NOT NULL
		);`,
	}

	for _, statement := range statements {
		result := db.Exec(statement)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
