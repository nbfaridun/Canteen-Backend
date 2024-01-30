package request

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/pkg/helpers"
)

type CreateSupplier struct {
	Name string `json:"name" validate:"required,min=1,max=50,alphanumunicode_and_space"`
}

type UpdateSupplier struct {
	Name string `json:"name" validate:"omitempty,min=1,max=50,alphanumunicode_and_space"`
}

func MapCreateSupplierToSupplier(input *CreateSupplier) *models.Supplier {
	return &models.Supplier{
		Name: input.Name,
	}
}

func MapUpdateSupplierToSupplier(input *UpdateSupplier) *models.Supplier {
	return &models.Supplier{
		Name: input.Name,
	}
}

type CreatePurchase struct {
	PurchaseDate         string                `json:"purchase_date" validate:"required,datetime=2006-01-02 15:04"`
	SupplierID           uint                  `json:"supplier_id" validate:"required,numeric"`
	TotalSum             float64               `json:"total_sum" validate:"required,numeric"`
	PurchasedIngredients []PurchasedIngredient `json:"ingredients"`
}

// todo добавить валидацию больше нуля
type PurchasedIngredient struct {
	ID             uint    `json:"id" validate:"required,numeric"`
	Name           string  `json:"name" validate:"required,min=1,max=50,alphanumunicode_and_space"`
	Amount         float64 `json:"amount" validate:"required,numeric"`
	Cost           float64 `json:"cost" validate:"required,numeric"`
	ExpirationDate string  `json:"expiration_date" validate:"required,datetime=2006-01-02"`
}

func MapCreatePurchaseToPurchase(input *CreatePurchase) *models.Purchase {
	purchase := &models.Purchase{
		PurchaseDate: helpers.ConvertStringToDate(input.PurchaseDate, "2006-01-02 15:04"),
		SupplierID:   input.SupplierID,
		TotalSum:     input.TotalSum,
	}

	for _, ingredient := range input.PurchasedIngredients {
		purchase.PurchasedIngredients = append(purchase.PurchasedIngredients, models.PurchasedIngredients{
			ID:             ingredient.ID,
			Name:           ingredient.Name,
			Amount:         ingredient.Amount,
			Cost:           ingredient.Cost,
			ExpirationDate: helpers.ConvertStringToDate(ingredient.ExpirationDate, "2006-01-02"),
		})
	}

	return purchase
}
