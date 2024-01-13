package request

import "Canteen-Backend/internal/models"

type CreateClientCategory struct {
	Name string `json:"name" validate:"required,min=1,max=20,alpha"`
}

type UpdateClientCategory struct {
	Name     string `json:"name" validate:"omitempty,min=1,max=20,alpha"`
	IsActive bool   `json:"is_active"`
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
