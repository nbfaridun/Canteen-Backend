package repository

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository/postgres"
	"gorm.io/gorm"
)

type User interface {
	GetAllUsers() (*[]models.User, error)
	CreateUser(user *models.User) (uint, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(ud uint) error
	GetRoleByID(id uint) (string, error)
}

type ClientCategory interface {
	CreateClientCategory(clientCategory *models.ClientCategory) (uint, error)
	GetAllClientCategories() (*[]models.ClientCategory, error)
	GetClientCategoryByID(id uint) (*models.ClientCategory, error)
	UpdateClientCategory(id uint, clientCategory *models.ClientCategory) error
	DeleteClientCategory(id uint) error
}

type Client interface {
	CreateClient(client *models.Client) (uint, error)
	GetAllClients() (*[]models.Client, error)
	GetClientByID(id uint) (*models.Client, error)
	UpdateClient(id uint, client *models.Client) error
	DeleteClient(id uint) error
}

type Repository struct {
	User
	ClientCategory
	Client
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:           postgres.NewUserPostgres(db),
		ClientCategory: postgres.NewClientCategoryPostgres(db),
		Client:         postgres.NewClientPostgres(db),
	}
}
