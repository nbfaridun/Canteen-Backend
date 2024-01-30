package request

import "Canteen-Backend/internal/models"

type CreateClient struct {
	Email            string  `json:"email" validate:"required,email"`
	FirstName        string  `json:"first_name" validate:"required,min=1,max=20,alpha"`
	LastName         string  `json:"last_name" validate:"required,min=1,max=20,alpha"`
	Age              uint    `json:"age" validate:"required,min=1,max=100"`
	Gender           string  `json:"gender" validate:"required,gender"`
	Balance          float32 `json:"balance" validate:"omitempty"`
	ClientCategoryID uint    `json:"client_category_id" validate:"required"`
}

type UpdateClient struct {
	Email            string  `json:"email" validate:"omitempty,email"`
	FirstName        string  `json:"first_name" validate:"omitempty,min=1,max=20,alpha"`
	LastName         string  `json:"last_name" validate:"omitempty,min=1,max=20,alpha"`
	Age              uint    `json:"age" validate:"omitempty,min=1,max=100"`
	Gender           string  `json:"gender" validate:"omitempty"`
	Balance          float32 `json:"balance" validate:"omitempty"`
	ClientCategoryID uint    `json:"client_category_id" validate:"omitempty"`
	IsActive         bool    `json:"is_active"`
}

type ModifyBalance struct {
	Difference float32 `json:"difference" validate:"required"`
}

type CreateClientCategory struct {
	Name string `json:"name" validate:"required,min=1,max=20,alpha"`
}

type UpdateClientCategory struct {
	Name     string `json:"name" validate:"omitempty,min=1,max=20,alpha"`
	IsActive bool   `json:"is_active"`
}

func MapCreateClientToClient(input *CreateClient) *models.Client {
	return &models.Client{
		Email:            input.Email,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		ClientCategoryID: input.ClientCategoryID,
		Balance:          input.Balance,
		Age:              input.Age,
		Gender:           input.Gender,
		IsActive:         true,
	}
}

func MapUpdateClientToClient(input *UpdateClient) *models.Client {
	return &models.Client{
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Age:              input.Age,
		Gender:           input.Gender,
		Email:            input.Email,
		ClientCategoryID: input.ClientCategoryID,
		Balance:          input.Balance,
		IsActive:         input.IsActive,
	}
}

func MapCreateClientCategoryToClientCategory(input *CreateClientCategory) *models.ClientCategory {
	return &models.ClientCategory{
		Name:     input.Name,
		IsActive: true,
	}
}

func MapUpdateClientCategoryToClientCategory(input *UpdateClientCategory) *models.ClientCategory {
	return &models.ClientCategory{
		Name:     input.Name,
		IsActive: input.IsActive,
	}
}
