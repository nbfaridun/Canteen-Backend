package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/customErr"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type ClientCategoryUseCase struct {
	repo repository.ClientCategory
}

func NewClientCategoryUseCase(repo repository.ClientCategory) *ClientCategoryUseCase {
	return &ClientCategoryUseCase{repo: repo}
}

func (cc *ClientCategoryUseCase) CreateClientCategory(clientCategory *models.ClientCategory) (uint, *customErr.CustomError) {
	id, err := cc.repo.CreateClientCategory(clientCategory)
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

func (cc *ClientCategoryUseCase) GetAllClientCategories() (*[]models.ClientCategory, *customErr.CustomError) {
	clientCategories, err := cc.repo.GetAllClientCategories()
	if err != nil {
		return nil, &customErr.CustomError{
			Error:      err,
			Message:    customErr.ServerError.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return clientCategories, nil
}

func (cc *ClientCategoryUseCase) GetClientCategoryByID(id uint) (*models.ClientCategory, *customErr.CustomError) {
	clientCategory, err := cc.repo.GetClientCategoryByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customErr.CustomError{
				Error:      err,
				Message:    customErr.ClientCategoryNotFound.Error(),
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

	return clientCategory, nil
}

func (cc *ClientCategoryUseCase) UpdateClientCategory(id uint, clientCategory *models.ClientCategory) *customErr.CustomError {
	err := cc.repo.UpdateClientCategory(id, clientCategory)
	if err != nil {
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

	return nil
}

func (cc *ClientCategoryUseCase) DeleteClientCategory(id uint) *customErr.CustomError {
	err := cc.repo.DeleteClientCategory(id)
	if err != nil {
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

	return nil
}
