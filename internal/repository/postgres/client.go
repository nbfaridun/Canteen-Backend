package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type ClientPostgres struct {
	db *gorm.DB
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (r *ClientPostgres) CreateClient(client *models.Client) (uint, error) {
	result := r.db.Table(constants.ClientTableName).Create(client)
	if result.Error != nil {
		return 0, result.Error
	}

	return client.ID, nil
}

func (r *ClientPostgres) GetAllClients() (*[]models.Client, error) {
	var clients []models.Client
	result := r.db.Table(constants.ClientTableName).Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clients, nil
}

func (r *ClientPostgres) GetAllClientsByCategoryID(clientCategoryID uint) (*[]models.Client, error) {
	var clients []models.Client
	result := r.db.Table(constants.ClientTableName).Find(&clients, "client_category_id = ?", clientCategoryID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clients, nil
}

func (r *ClientPostgres) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	result := r.db.Table(constants.ClientTableName).First(&client, "client_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &client, nil
}

func (r *ClientPostgres) UpdateClient(client *models.Client) error {
	result := r.db.Table(constants.ClientTableName).Model(&models.Client{}).Where("client_id = ?", client.ID).Updates(client)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ClientPostgres) DeleteClient(id uint) error {
	result := r.db.Table(constants.ClientTableName).Delete(&models.Client{}, "client_id = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ClientPostgres) CreateClientCategory(clientCategory *models.ClientCategory) (uint, error) {
	result := r.db.Table(constants.ClientCategoryTableName).Create(clientCategory)
	if result.Error != nil {
		return 0, result.Error
	}

	return clientCategory.ID, nil
}

func (r *ClientPostgres) GetAllClientCategories() (*[]models.ClientCategory, error) {
	var clientCategories []models.ClientCategory
	result := r.db.Table(constants.ClientCategoryTableName).Find(&clientCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clientCategories, nil
}

func (r *ClientPostgres) GetClientCategoryByID(id uint) (*models.ClientCategory, error) {
	var clientCategory models.ClientCategory
	result := r.db.Table(constants.ClientCategoryTableName).First(&clientCategory, "client_category_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clientCategory, nil
}

func (r *ClientPostgres) GetClientCategoryByName(clientCategoryName string) (*models.ClientCategory, error) {
	var clientCategory models.ClientCategory
	result := r.db.Table(constants.ClientCategoryTableName).First(&clientCategory, "name = ?", clientCategoryName)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clientCategory, nil
}

func (r *ClientPostgres) UpdateClientCategory(clientCategory *models.ClientCategory) error {
	result := r.db.Table(constants.ClientCategoryTableName).Model(&models.ClientCategory{}).Where("client_category_id = ?", clientCategory.ID).Updates(clientCategory)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ClientPostgres) DeleteClientCategory(id uint) error {
	result := r.db.Table(constants.ClientCategoryTableName).Delete(&models.ClientCategory{}, "client_category_id = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
