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
	UpdateUser(user *models.User) *customErr.CustomError
	DeleteUser(id uint) *customErr.CustomError
	SignIn(userInput *models.User) (*models.Token, *customErr.CustomError)
	RefreshTokens(refreshToken string) (*models.Token, *customErr.CustomError)
	SignOut(refreshToken string) *customErr.CustomError
}

type Client interface {
	CreateClient(client *models.Client) (uint, *customErr.CustomError)
	GetAllClients(clientCategoryName string) (*[]models.Client, *customErr.CustomError)
	GetClientByID(id uint) (*models.Client, *customErr.CustomError)
	UpdateClient(client *models.Client) *customErr.CustomError
	DeleteClient(id uint) *customErr.CustomError
	ModifyBalanceByClientID(id uint, difference float32) *customErr.CustomError

	CreateClientCategory(clientCategory *models.ClientCategory) (uint, *customErr.CustomError)
	GetAllClientCategories() (*[]models.ClientCategory, *customErr.CustomError)
	GetClientCategoryByID(id uint) (*models.ClientCategory, *customErr.CustomError)
	UpdateClientCategory(clientCategory *models.ClientCategory) *customErr.CustomError
	DeleteClientCategory(id uint) *customErr.CustomError
}

type Ingredient interface {
	CreateIngredientCategory(ingredientCategory *models.IngredientCategory) (uint, *customErr.CustomError)
	GetAllIngredientCategories() (*[]models.IngredientCategory, *customErr.CustomError)
	GetIngredientCategoryByID(id uint) (*models.IngredientCategory, *customErr.CustomError)
	UpdateIngredientCategory(ingredientCategory *models.IngredientCategory) *customErr.CustomError
	DeleteIngredientCategory(id uint) *customErr.CustomError

	CreateIngredient(ingredient *models.Ingredient) (uint, *customErr.CustomError)
	GetAllIngredients() (*[]models.Ingredient, *customErr.CustomError)
	GetIngredientByID(id uint) (*models.Ingredient, *customErr.CustomError)
	UpdateIngredient(ingredient *models.Ingredient) *customErr.CustomError
	DeleteIngredient(id uint) *customErr.CustomError
}

type Purchase interface {
	CreateSupplier(supplier *models.Supplier) (uint, *customErr.CustomError)
	GetAllSuppliers() (*[]models.Supplier, *customErr.CustomError)
	GetSupplierByID(id uint) (*models.Supplier, *customErr.CustomError)
	UpdateSupplier(supplier *models.Supplier) *customErr.CustomError
	DeleteSupplier(id uint) *customErr.CustomError

	CreatePurchase(purchase *models.Purchase) (uint, *customErr.CustomError)
}

type UseCase struct {
	User
	Client
	Ingredient
	Purchase
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:       NewUserUseCase(repo.User, repo.Session),
		Client:     NewClientUseCase(repo.Client),
		Ingredient: NewIngredientUseCase(repo.Ingredient),
		Purchase:   NewPurchaseUseCase(repo.Purchase, repo.Ingredient),
	}
}
