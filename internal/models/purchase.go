package models

import "time"

type Supplier struct {
	ID   uint   `gorm:"column:supplier_id"`
	Name string `gorm:"column:name"`
}

type Purchase struct {
	ID                   uint                   `gorm:"column:purchase_id"`
	PurchaseDate         time.Time              `gorm:"column:purchase_date"`
	SupplierID           uint                   `gorm:"column:supplier_id"`
	TotalSum             float64                `gorm:"column:total_sum"`
	PurchasedIngredients []PurchasedIngredients `gorm:"-"`
}

type PurchasedIngredients struct {
	ID             uint
	Name           string
	Amount         float64
	Cost           float64
	ExpirationDate time.Time
}

// todo add here amount, cost, current_unit_price
type PurchasesIngredients struct {
	PurchaseID       uint    `gorm:"column:purchase_id"`
	IngredientID     uint    `gorm:"column:ingredient_id"`
	Amount           float64 `gorm:"column:amount"`
	Cost             float64 `gorm:"column:cost"`
	CurrentUnitPrice float64 `gorm:"column:current_unit_price"`
}
