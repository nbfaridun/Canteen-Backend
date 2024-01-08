package models

import (
	"time"
)

var ClientTableName = "client"
var ClientCategoryTableName = "client_category"

type Client struct {
	ID               uint      `gorm:"column:client_id;primaryKey" json:"id"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	FirstName        string    `gorm:"column:first_name" json:"first_name"`
	LastName         string    `gorm:"column:last_name" json:"last_name"`
	Age              uint      `gorm:"column:age" json:"age"`
	Gender           string    `gorm:"column:gender" json:"gender"`
	Email            string    `gorm:"column:email" json:"email"`
	ClientCategoryID uint      `gorm:"column:client_category_id" json:"client_category_id"`
	Balance          float32   `gorm:"column:balance" json:"balance"`
	IsActive         bool      `gorm:"column:is_active" json:"is_active"`
}

type CreateClientInput struct {
	FirstName        string  `json:"first_name" binding:"required"`
	LastName         string  `json:"last_name" binding:"required"`
	Age              uint    `json:"age" binding:"required"`
	Gender           string  `json:"gender" binding:"required"`
	Email            string  `json:"email"`
	ClientCategoryID uint    `json:"client_category_id" binding:"required"`
	Balance          float32 `json:"balance"`
}

type UpdateClientInput struct {
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	Age              uint    `json:"age"`
	Gender           string  `json:"gender"`
	Email            string  `json:"email"`
	ClientCategoryID uint    `json:"client_category_id"`
	Balance          float32 `json:"balance"`
	IsActive         bool    `json:"is_active"`
}

type ClientCategory struct {
	ID        uint      `gorm:"column:client_category_id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
}
