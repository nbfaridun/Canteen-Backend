package migrations

import (
	"Canteen-Backend/pkg/helpers"
	"gorm.io/gorm"
	"os"
)

func RunSeeds(db *gorm.DB) error {
	if err := seedIfNotExists(db, "user_role", "name", "admin"); err != nil {
		return err
	}
	if err := seedIfNotExists(db, "user_role", "name", "receptionist"); err != nil {
		return err
	}
	if err := seedIfNotExists(db, "user_role", "name", "cashier"); err != nil {
		return err
	}

	if err := seedIfNotExists(db, "client_category", "name", "students"); err != nil {
		return err
	}
	if err := seedIfNotExists(db, "client_category", "name", "faculty"); err != nil {
		return err
	}
	if err := seedIfNotExists(db, "client_category", "name", "staff"); err != nil {
		return err
	}

	if err := seedIfNotExists(db, "ingredient_category", "name", "vegetables"); err != nil {
		return err
	}

	if err := seedIfNotExists(db, "supplier", "name", "supplier1"); err != nil {
		return err
	}

	if err := createAdmin(db); err != nil {
		return err
	}

	return nil
}

func createAdmin(db *gorm.DB) error {
	//checking if admin exists
	var count int64
	result := db.Table("user").Where("username = ? OR email = ?", os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_EMAIL")).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count == 0 {
		createStatement := `INSERT INTO "user" (user_role_id, username, email, first_name, last_name, password) VALUES (?, ?, ?, ?, ?, ?);`
		adminPassword, err := helpers.HashPassword(os.Getenv("ADMIN_PASSWORD"))
		if err != nil {
			return err
		}

		result := db.Exec(createStatement, os.Getenv("ADMIN_ROLE_ID"), os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_FIRST_NAME"), os.Getenv("ADMIN_LAST_NAME"), adminPassword)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func seedIfNotExists(db *gorm.DB, tableName, columnName, value string) error {
	var count int64
	result := db.Table(tableName).Where(columnName+" = ?", value).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count == 0 {
		// Record does not exist, so insert it
		insertStatement := "INSERT INTO " + tableName + " (" + columnName + ") VALUES (?);"
		result := db.Exec(insertStatement, value)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
