package request

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/pkg/helpers"
)

type CreateIngredient struct {
	Name                 string `json:"name" validate:"required,min=1,max=50,alphanumunicode_and_space"`
	IngredientCategoryID uint   `json:"ingredient_category_id" validate:"required,number"`
	Unit                 string `json:"unit" validate:"required,alphaunicode"`
}

// todo добавить валидацию больше нуля
type UpdateIngredient struct {
	Name                 string  `json:"name" validate:"omitempty,min=1,max=50,alphanumunicode_and_space"`
	IngredientCategoryID uint    `json:"ingredient_category_id" validate:"omitempty,number"`
	Unit                 string  `json:"unit" validate:"omitempty,alphaunicode"`
	Quantity             float64 `json:"quantity" validate:"omitempty,number"`
	UnitPrice            float64 `json:"unit_price" validate:"omitempty,number"`
	LackLimit            float64 `json:"lack_limit" validate:"omitempty,number"`
	PurchaseDate         string  `json:"purchase_date" validate:"omitempty,datetime=2006-01-02 15:04"`
	ExpirationDate       string  `json:"expiration_date" validate:"omitempty,datetime=2006-01-02"`
}

func MapCreateIngredientToIngredient(input *CreateIngredient) *models.Ingredient {
	return &models.Ingredient{
		Name:                 input.Name,
		IngredientCategoryID: input.IngredientCategoryID,
		Unit:                 input.Unit,
	}
}

func MapUpdateIngredientToIngredient(input *UpdateIngredient) *models.Ingredient {
	return &models.Ingredient{
		Name:                 input.Name,
		IngredientCategoryID: input.IngredientCategoryID,
		Unit:                 input.Unit,
		Quantity:             input.Quantity,
		UnitPrice:            input.UnitPrice,
		LackLimit:            input.LackLimit,
		PurchaseDate:         helpers.ConvertStringToDate(input.PurchaseDate, "2006-01-02 15:04"),
		ExpirationDate:       helpers.ConvertStringToDate(input.ExpirationDate, "2006-01-02"),
	}
}

type CreateIngredientCategory struct {
	Name string `json:"name" validate:"required,min=1,max=50,alphanumunicode_and_space"`
}

type UpdateIngredientCategory struct {
	Name string `json:"name" validate:"omitempty,min=1,max=50,alphanumunicode_and_space"`
}

func MapCreateIngredientCategoryToIngredientCategory(input *CreateIngredientCategory) *models.IngredientCategory {
	return &models.IngredientCategory{
		Name: input.Name,
	}
}

func MapUpdateIngredientCategoryToIngredientCategory(input *UpdateIngredientCategory) *models.IngredientCategory {
	return &models.IngredientCategory{
		Name: input.Name,
	}
}
