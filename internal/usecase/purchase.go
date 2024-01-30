package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/customErr"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type PurchaseUseCase struct {
	repoPurchase   repository.Purchase
	repoIngredient repository.Ingredient
}

func NewPurchaseUseCase(repoPurchase repository.Purchase, repoIngredient repository.Ingredient) *PurchaseUseCase {
	return &PurchaseUseCase{repoPurchase: repoPurchase, repoIngredient: repoIngredient}
}

func (u *PurchaseUseCase) CreateSupplier(supplier *models.Supplier) (uint, *customErr.CustomError) {
	id, err := u.repoPurchase.CreateSupplier(supplier)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.SupplierAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *PurchaseUseCase) GetAllSuppliers() (*[]models.Supplier, *customErr.CustomError) {
	suppliers, err := u.repoPurchase.GetAllSuppliers()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return suppliers, nil
}

func (u *PurchaseUseCase) GetSupplierByID(id uint) (*models.Supplier, *customErr.CustomError) {
	supplier, err := u.repoPurchase.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.SupplierNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return supplier, nil
}

func (u *PurchaseUseCase) UpdateSupplier(supplier *models.Supplier) *customErr.CustomError {

	err := u.repoPurchase.UpdateSupplier(supplier)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return customErr.NewCustomError(err, customErr.SupplierAlreadyExists.Error(), http.StatusConflict)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.SupplierNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *PurchaseUseCase) DeleteSupplier(id uint) *customErr.CustomError {
	err := u.repoPurchase.DeleteSupplier(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.SupplierNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *PurchaseUseCase) CreatePurchase(purchase *models.Purchase) (uint, *customErr.CustomError) {
	if _, err := u.repoPurchase.GetSupplierByID(purchase.SupplierID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErr.NewCustomError(err, customErr.SupplierNotFound.Error(), http.StatusNotFound)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	id, err := u.repoPurchase.CreatePurchase(purchase)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.PurchaseAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	// here we update ingredients in storage
	for _, purchasedIngredient := range purchase.PurchasedIngredients {
		ingredientInStorage, err := u.repoIngredient.GetIngredientByID(purchasedIngredient.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
			} else {
				return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}

		ingredient := &models.Ingredient{
			ID:             purchasedIngredient.ID,
			Quantity:       ingredientInStorage.Quantity + purchasedIngredient.Amount,
			ExpirationDate: purchasedIngredient.ExpirationDate,
			PurchaseDate:   purchase.PurchaseDate,
			UnitPrice:      ingredientInStorage.UnitPrice + purchasedIngredient.Cost/purchasedIngredient.Amount,
		}

		err = u.repoIngredient.UpdateIngredient(ingredient)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
			} else {
				return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}
	}

	// here we create purchases_ingredients
	var purchasesIngredientsSlice []models.PurchasesIngredients
	for _, ingredient := range purchase.PurchasedIngredients {
		var purchasesIngredients models.PurchasesIngredients
		purchasesIngredients.PurchaseID = id
		purchasesIngredients.IngredientID = ingredient.ID
		purchasesIngredients.Amount = ingredient.Amount
		purchasesIngredients.Cost = ingredient.Cost
		purchasesIngredients.CurrentUnitPrice = ingredient.Cost / ingredient.Amount
		purchasesIngredientsSlice = append(purchasesIngredientsSlice, purchasesIngredients)
	}

	err = u.repoPurchase.CreatePurchasesIngredients(&purchasesIngredientsSlice)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}
