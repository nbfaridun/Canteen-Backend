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
		`CREATE TABLE IF NOT EXISTS ingredient_category (
    		ingredient_category_id SERIAL PRIMARY KEY,
    		name VARCHAR(100) UNIQUE NOT NULL,
    		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    	);`,
		`CREATE TABLE IF NOT EXISTS ingredient (
		ingredient_id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL,
		ingredient_category_id INT NOT NULL REFERENCES ingredient_category(ingredient_category_id) ON DELETE CASCADE,
		unit VARCHAR(20) NOT NULL,
		quantity FLOAT,
		unit_price FLOAT,
		lack_limit FLOAT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		purchase_date TIMESTAMP,
		expiration_date TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS supplier (
    		supplier_id SERIAL PRIMARY KEY,
    		name VARCHAR(100) UNIQUE NOT NULL
        );`,
		`CREATE TABLE IF NOT EXISTS purchase (
    	purchase_id SERIAL PRIMARY KEY,
    	purchase_date TIMESTAMP NOT NULL,
    	supplier_id INT NOT NULL REFERENCES supplier(supplier_id) ON DELETE CASCADE,
    	total_sum FLOAT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS purchases_ingredients (
    	purchase_id INT NOT NULL REFERENCES purchase(purchase_id) ON DELETE CASCADE,
    	ingredient_id INT NOT NULL REFERENCES ingredient(ingredient_id) ON DELETE CASCADE,
    	amount FLOAT NOT NULL,
    	cost FLOAT NOT NULL,
    	current_unit_price FLOAT NOT NULL
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
