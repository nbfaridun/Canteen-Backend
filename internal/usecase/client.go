package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/customErr"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type ClientUseCase struct {
	repoClient repository.Client
}

func NewClientUseCase(repoClient repository.Client) *ClientUseCase {
	return &ClientUseCase{repoClient: repoClient}
}

func (u *ClientUseCase) CreateClient(client *models.Client) (uint, *customErr.CustomError) {
	if _, err := u.repoClient.GetClientCategoryByID(client.ClientCategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}
	id, err := u.repoClient.CreateClient(client)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.EmailAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *ClientUseCase) GetAllClients(clientCategoryName string) (*[]models.Client, *customErr.CustomError) {
	if clientCategoryName != "" {
		clientCategory, err := u.repoClient.GetClientCategoryByName(clientCategoryName)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
			} else {
				return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}

		clients, err := u.repoClient.GetAllClientsByCategoryID(clientCategory.ID)
		if err != nil {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}

		return clients, nil

	}

	clients, err := u.repoClient.GetAllClients()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return clients, nil
}

func (u *ClientUseCase) GetClientByID(id uint) (*models.Client, *customErr.CustomError) {
	client, err := u.repoClient.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.ClientNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return client, nil
}

func (u *ClientUseCase) UpdateClient(client *models.Client) *customErr.CustomError {

	if client.ClientCategoryID != 0 {
		if _, err := u.repoClient.GetClientCategoryByID(client.ClientCategoryID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
			} else {
				return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}
	}

	if err := u.repoClient.UpdateClient(client); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *ClientUseCase) DeleteClient(id uint) *customErr.CustomError {
	if err := u.repoClient.DeleteClient(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *ClientUseCase) ModifyBalanceByClientID(id uint, difference float32) *customErr.CustomError {
	user, err := u.repoClient.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.ClientNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	user.Balance += difference

	if err := u.repoClient.UpdateClient(user); err != nil {
		return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (u *ClientUseCase) CreateClientCategory(clientCategory *models.ClientCategory) (uint, *customErr.CustomError) {
	id, err := u.repoClient.CreateClientCategory(clientCategory)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.ClientCategoryAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *ClientUseCase) GetAllClientCategories() (*[]models.ClientCategory, *customErr.CustomError) {
	clientCategories, err := u.repoClient.GetAllClientCategories()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return clientCategories, nil
}

func (u *ClientUseCase) GetClientCategoryByID(id uint) (*models.ClientCategory, *customErr.CustomError) {
	clientCategory, err := u.repoClient.GetClientCategoryByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return clientCategory, nil
}

func (u *ClientUseCase) UpdateClientCategory(clientCategory *models.ClientCategory) *customErr.CustomError {
	err := u.repoClient.UpdateClientCategory(clientCategory)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *ClientUseCase) DeleteClientCategory(id uint) *customErr.CustomError {
	err := u.repoClient.DeleteClientCategory(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.ClientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}
