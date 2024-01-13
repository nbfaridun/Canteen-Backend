package models

import (
	"time"
)

type Client struct {
	ID               uint      `gorm:"column:client_id;primaryKey"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at"`
	FirstName        string    `gorm:"column:first_name"`
	LastName         string    `gorm:"column:last_name"`
	Age              uint      `gorm:"column:age"`
	Gender           string    `gorm:"column:gender"`
	Email            string    `gorm:"column:email"`
	ClientCategoryID uint      `gorm:"column:client_category_id"`
	Balance          float32   `gorm:"column:balance"`
	IsActive         bool      `gorm:"column:is_active"`
}

type ClientCategory struct {
	ID        uint      `gorm:"column:client_category_id;primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	IsActive  bool      `gorm:"column:is_active"`
}
