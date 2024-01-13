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
	repoClient         repository.Client
	repoClientCategory repository.ClientCategory
}

func NewClientUseCase(repoClient repository.Client, repoClientCategory repository.ClientCategory) *ClientUseCase {
	return &ClientUseCase{repoClient: repoClient, repoClientCategory: repoClientCategory}
}

func (u *ClientUseCase) CreateClient(client *models.Client) (uint, *customErr.CustomError) {
	if _, err := u.repoClientCategory.GetClientCategoryByID(client.ClientCategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientCategoryNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return 0, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	id, err := u.repoClient.CreateClient(client)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, &customErr.CustomError{
				Error:      err,
				Message:    customErr.EmailAlreadyExists.Error(),
				StatusCode: http.StatusConflict,
			}
		} else {
			return 0, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	return id, nil
}

func (u *ClientUseCase) GetAllClients() (*[]models.Client, *customErr.CustomError) {
	clients, err := u.repoClient.GetAllClients()
	if err != nil {
		return nil, &customErr.CustomError{
			Error:      err,
			Message:    customErr.ServerError.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return clients, nil
}

func (u *ClientUseCase) GetClientByID(id uint) (*models.Client, *customErr.CustomError) {
	client, err := u.repoClient.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	return client, nil
}

func (u *ClientUseCase) UpdateClient(id uint, client *models.Client) *customErr.CustomError {
	if _, err := u.repoClient.GetClientByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	if client.ClientCategoryID != 0 {
		if _, err := u.repoClientCategory.GetClientCategoryByID(client.ClientCategoryID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &customErr.CustomError{
					Error:      err,
					Message:    customErr.ClientCategoryNotFound.Error(),
					StatusCode: http.StatusNotFound,
				}
			} else {
				return &customErr.CustomError{
					Error:      err,
					Message:    customErr.ServerError.Error(),
					StatusCode: http.StatusInternalServerError,
				}
			}
		}
	}

	if err := u.repoClient.UpdateClient(id, client); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	return nil
}

func (u *ClientUseCase) DeleteClient(id uint) *customErr.CustomError {
	if err := u.repoClient.DeleteClient(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	return nil
}

func (u *ClientUseCase) ModifyBalanceByClientID(id uint, difference float32) *customErr.CustomError {
	user, err := u.repoClient.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientNotFound.Error(),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return &customErr.CustomError{
				Error:      err,
				Message:    customErr.ServerError.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	user.Balance += difference

	if err := u.repoClient.UpdateClient(id, user); err != nil {
		return &customErr.CustomError{
			Error:      err,
			Message:    customErr.ServerError.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
