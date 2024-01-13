package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type ClientCategoryPostgres struct {
	db *gorm.DB
}

func NewClientCategoryPostgres(db *gorm.DB) *ClientCategoryPostgres {
	return &ClientCategoryPostgres{db: db}
}

func (r *ClientCategoryPostgres) CreateClientCategory(clientCategory *models.ClientCategory) (uint, error) {
	result := r.db.Table(constants.ClientCategoryTableName).Create(clientCategory)
	if result.Error != nil {
		return 0, result.Error
	}

	return clientCategory.ID, nil
}

func (r *ClientCategoryPostgres) GetAllClientCategories() (*[]models.ClientCategory, error) {
	var clientCategories []models.ClientCategory
	result := r.db.Table(constants.ClientCategoryTableName).Find(&clientCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clientCategories, nil
}

func (r *ClientCategoryPostgres) GetClientCategoryByID(id uint) (*models.ClientCategory, error) {
	var clientCategory models.ClientCategory
	result := r.db.Table(constants.ClientCategoryTableName).First(&clientCategory, "client_category_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clientCategory, nil
}

func (r *ClientCategoryPostgres) UpdateClientCategory(id uint, clientCategory *models.ClientCategory) error {
	result := r.db.Table(constants.ClientCategoryTableName).Model(&models.ClientCategory{}).Where("client_category_id = ?", id).Updates(clientCategory)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ClientCategoryPostgres) DeleteClientCategory(id uint) error {
	result := r.db.Table(constants.ClientCategoryTableName).Delete(&models.ClientCategory{}, "client_category_id = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
