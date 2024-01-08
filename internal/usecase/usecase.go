package usecase

import (
	"Canteen-Backend/customErr"
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
)

type User interface {
	CreateUser(user *models.User) (uint, *customErr.CustomError)
	GetAllUsers() (*[]models.User, *customErr.CustomError)
	GetUserByID(id uint) (*models.User, *customErr.CustomError)
	UpdateUser(id uint, user *models.User) *customErr.CustomError
	DeleteUser(id uint) *customErr.CustomError
}

type Client interface {
	CreateClient(client *models.Client) (uint, *customErr.CustomError)
	GetAllClients() (*[]models.Client, *customErr.CustomError)
	GetClientByID(id uint) (*models.Client, *customErr.CustomError)
	UpdateClient(id uint, client *models.Client) *customErr.CustomError
	DeleteClient(id uint) *customErr.CustomError
}

type UseCase struct {
	User
	Client
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:   NewUserUseCase(repo.User),
		Client: NewClientUseCase(repo.Client, repo.ClientCategory),
	}
}
