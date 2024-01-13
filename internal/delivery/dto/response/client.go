package response

import "Canteen-Backend/internal/models"

type GetClient struct {
	ID               uint    `json:"id"`
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	Age              uint    `json:"age"`
	Gender           string  `json:"gender"`
	Email            string  `json:"email"`
	ClientCategoryID uint    `json:"client_category_id"`
	Balance          float32 `json:"balance"`
	IsActive         bool    `json:"is_active"`
}

func MapClientToGetClient(client *models.Client) *GetClient {
	return &GetClient{
		ID:               client.ID,
		FirstName:        client.FirstName,
		LastName:         client.LastName,
		Age:              client.Age,
		Gender:           client.Gender,
		Email:            client.Email,
		ClientCategoryID: client.ClientCategoryID,
		Balance:          client.Balance,
		IsActive:         client.IsActive,
	}
}
