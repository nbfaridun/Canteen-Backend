package response

import "Canteen-Backend/internal/models"

type GetIngredient struct {
	ID                   uint    `json:"id"`
	Name                 string  `json:"name"`
	IngredientCategoryID uint    `json:"ingredient_category_id"`
	Unit                 string  `json:"unit"`
	Quantity             float64 `json:"quantity"`
	UnitPrice            float64 `json:"unit_price"`
	LackLimit            float64 `json:"lack_limit"`
	PurchaseDate         string  `json:"purchase_date"`
	ExpirationDate       string  `json:"expiration_date"`
}

func MapIngredientToGetIngredient(ingredient *models.Ingredient) *GetIngredient {
	return &GetIngredient{
		ID:                   ingredient.ID,
		Name:                 ingredient.Name,
		IngredientCategoryID: ingredient.IngredientCategoryID,
		Unit:                 ingredient.Unit,
		Quantity:             ingredient.Quantity,
		UnitPrice:            ingredient.UnitPrice,
		LackLimit:            ingredient.LackLimit,
		PurchaseDate:         ingredient.PurchaseDate.Format("2006-01-02 15:04"),
		ExpirationDate:       ingredient.ExpirationDate.Format("2006-01-02"),
	}
}

type GetIngredientCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func MapIngredientCategoryToGetIngredientCategory(ingredientCategory *models.IngredientCategory) *GetIngredientCategory {
	return &GetIngredientCategory{
		ID:   ingredientCategory.ID,
		Name: ingredientCategory.Name,
	}
}
