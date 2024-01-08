package models

import (
	"time"
)

var UserTableName = "user"
var RoleTableName = "role"

type User struct {
	ID        uint      `gorm:"column:user_id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	RoleID    uint      `gorm:"column:role_id" json:"role_id"`
	Username  string    `gorm:"column:username;unique" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email;unique" json:"email"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
}

type Role struct {
	ID   uint   `gorm:"column:role_id;primaryKey" json:"id"`
	Name string `gorm:"column:name;unique" json:"name"`
}

type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	RoleID    uint   `json:"role_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UpdateUserInput struct {
	Username  string `json:"username"`
	RoleID    uint   `json:"role_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"`
}
