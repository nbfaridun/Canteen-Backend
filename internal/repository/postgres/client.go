package postgres

import (
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
	result := r.db.Table(models.ClientTableName).Create(client)
	if result.Error != nil {
		return 0, result.Error
	}

	return client.ID, nil
}

func (r *ClientPostgres) GetAllClients() (*[]models.Client, error) {
	var clients []models.Client
	result := r.db.Table(models.ClientTableName).Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}

	return &clients, nil
}

func (r *ClientPostgres) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	result := r.db.Table(models.ClientTableName).First(&client, "client_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &client, nil
}

func (r *ClientPostgres) UpdateClient(id uint, client *models.Client) error {
	result := r.db.Table(models.ClientTableName).Model(&models.Client{}).Where("client_id = ?", id).Updates(client)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ClientPostgres) DeleteClient(id uint) error {
	result := r.db.Table(models.ClientTableName).Delete(&models.Client{}, "client_id = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
