package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/customErr"
)

type User interface {
	CreateUser(user *models.User) (uint, *customErr.CustomError)
	GetAllUsers() (*[]models.User, *customErr.CustomError)
	GetUserByID(id uint) (*models.User, *customErr.CustomError)
	UpdateUser(id uint, user *models.User) *customErr.CustomError
	DeleteUser(id uint) *customErr.CustomError
	SignIn(userInput *models.User) (*models.Token, *customErr.CustomError)
	RefreshTokens(refreshToken string) (*models.Token, *customErr.CustomError)
	SignOut(refreshToken string) *customErr.CustomError
}

type Client interface {
	CreateClient(client *models.Client) (uint, *customErr.CustomError)
	GetAllClients() (*[]models.Client, *customErr.CustomError)
	GetClientByID(id uint) (*models.Client, *customErr.CustomError)
	UpdateClient(id uint, client *models.Client) *customErr.CustomError
	DeleteClient(id uint) *customErr.CustomError
	ModifyBalanceByClientID(id uint, difference float32) *customErr.CustomError
}

type ClientCategory interface {
	CreateClientCategory(clientCategory *models.ClientCategory) (uint, *customErr.CustomError)
	GetAllClientCategories() (*[]models.ClientCategory, *customErr.CustomError)
	GetClientCategoryByID(id uint) (*models.ClientCategory, *customErr.CustomError)
	UpdateClientCategory(id uint, clientCategory *models.ClientCategory) *customErr.CustomError
	DeleteClientCategory(id uint) *customErr.CustomError
}

type UseCase struct {
	User
	Client
	ClientCategory
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:           NewUserUseCase(repo.User, repo.Session),
		Client:         NewClientUseCase(repo.Client, repo.ClientCategory),
		ClientCategory: NewClientCategoryUseCase(repo.ClientCategory),
	}
}
