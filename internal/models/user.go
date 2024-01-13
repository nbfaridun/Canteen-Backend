package models

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"column:user_id;primaryKey"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
	UserRoleID uint      `gorm:"column:user_role_id"`
	Username   string    `gorm:"column:username;unique"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email;unique"`
	FirstName  string    `gorm:"column:first_name"`
	LastName   string    `gorm:"column:last_name"`
	IsActive   bool      `gorm:"column:is_active"`
}

type UserRole struct {
	ID   uint   `gorm:"column:user_role_id;primaryKey" json:"id"`
	Name string `gorm:"column:name;unique" json:"name"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Session struct {
	ID           uint      `gorm:"column:session_id;primaryKey"`
	UserID       uint      `gorm:"column:user_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
}
