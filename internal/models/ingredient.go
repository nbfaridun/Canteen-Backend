package models

import "time"

type Ingredient struct {
	ID                   uint      `gorm:"column:ingredient_id"`
	Name                 string    `gorm:"column:name"`
	IngredientCategoryID uint      `gorm:"column:ingredient_category_id"`
	Unit                 string    `gorm:"column:unit"`
	Quantity             float64   `gorm:"column:quantity"`
	UnitPrice            float64   `gorm:"column:unit_price"`
	LackLimit            float64   `gorm:"column:lack_limit"`
	PurchaseDate         time.Time `gorm:"column:purchase_date"`
	ExpirationDate       time.Time `gorm:"column:expiration_date"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
}

type IngredientCategory struct {
	ID        uint      `gorm:"column:ingredient_category_id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
