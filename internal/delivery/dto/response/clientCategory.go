package response

import "Canteen-Backend/internal/models"

type GetClientCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func MapClientCategoryToGetClientCategory(clientCategory *models.ClientCategory) *GetClientCategory {
	return &GetClientCategory{
		ID:   clientCategory.ID,
		Name: clientCategory.Name,
	}
}
