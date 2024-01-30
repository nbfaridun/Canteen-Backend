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
	UpdateUser(user *models.User) error
	DeleteUser(ud uint) error
	GetRoleByID(id uint) (string, error)
	GetUserByUsername(username string) (*models.User, error)
}

type Session interface {
	CreateSession(session *models.Session) error
	DeleteSession(session *models.Session) error
	GetSessionByRefreshToken(refreshToken string) (*models.Session, error)
}

type Client interface {
	CreateClient(client *models.Client) (uint, error)
	GetAllClients() (*[]models.Client, error)
	GetAllClientsByCategoryID(clientCategoryID uint) (*[]models.Client, error)
	GetClientByID(id uint) (*models.Client, error)
	UpdateClient(client *models.Client) error
	DeleteClient(id uint) error

	CreateClientCategory(clientCategory *models.ClientCategory) (uint, error)
	GetAllClientCategories() (*[]models.ClientCategory, error)
	GetClientCategoryByID(id uint) (*models.ClientCategory, error)
	GetClientCategoryByName(clientCategoryName string) (*models.ClientCategory, error)
	UpdateClientCategory(clientCategory *models.ClientCategory) error
	DeleteClientCategory(id uint) error
}

type Ingredient interface {
	CreateIngredientCategory(ingredientCategory *models.IngredientCategory) (uint, error)
	GetAllIngredientCategories() (*[]models.IngredientCategory, error)
	GetIngredientCategoryByID(id uint) (*models.IngredientCategory, error)
	UpdateIngredientCategory(ingredientCategory *models.IngredientCategory) error
	DeleteIngredientCategory(id uint) error

	CreateIngredient(ingredient *models.Ingredient) (uint, error)
	GetAllIngredients() (*[]models.Ingredient, error)
	GetIngredientByID(id uint) (*models.Ingredient, error)
	UpdateIngredient(ingredient *models.Ingredient) error
	DeleteIngredient(id uint) error
}

type Purchase interface {
	CreateSupplier(supplier *models.Supplier) (uint, error)
	GetAllSuppliers() (*[]models.Supplier, error)
	GetSupplierByID(id uint) (*models.Supplier, error)
	UpdateSupplier(supplier *models.Supplier) error
	DeleteSupplier(id uint) error

	CreatePurchase(purchase *models.Purchase) (uint, error)
	CreatePurchasesIngredients(purchasesIngredients *[]models.PurchasesIngredients) error
}

type Repository struct {
	User
	Client
	Session
	Ingredient
	Purchase
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       postgres.NewUserPostgres(db),
		Client:     postgres.NewClientPostgres(db),
		Session:    postgres.NewSessionPostgres(db),
		Ingredient: postgres.NewIngredientPostgres(db),
		Purchase:   postgres.NewPurchasePostgres(db),
	}
}
