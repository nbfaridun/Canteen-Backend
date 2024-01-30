package usecase

import (
	"Canteen-Backend/internal/models"
	"Canteen-Backend/internal/repository"
	"Canteen-Backend/pkg/customErr"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type IngredientUseCase struct {
	repoIngredient repository.Ingredient
}

func NewIngredientUseCase(repoIngredient repository.Ingredient) *IngredientUseCase {
	return &IngredientUseCase{repoIngredient: repoIngredient}
}

func (u *IngredientUseCase) CreateIngredientCategory(ingredientCategory *models.IngredientCategory) (uint, *customErr.CustomError) {
	id, err := u.repoIngredient.CreateIngredientCategory(ingredientCategory)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.IngredientCategoryAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *IngredientUseCase) GetAllIngredientCategories() (*[]models.IngredientCategory, *customErr.CustomError) {
	ingredientCategories, err := u.repoIngredient.GetAllIngredientCategories()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return ingredientCategories, nil
}

func (u *IngredientUseCase) GetIngredientCategoryByID(id uint) (*models.IngredientCategory, *customErr.CustomError) {
	ingredientCategory, err := u.repoIngredient.GetIngredientCategoryByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.IngredientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return ingredientCategory, nil
}

func (u *IngredientUseCase) UpdateIngredientCategory(ingredientCategory *models.IngredientCategory) *customErr.CustomError {
	ingredientCategory.UpdatedAt = time.Now()

	if err := u.repoIngredient.UpdateIngredientCategory(ingredientCategory); err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return customErr.NewCustomError(err, customErr.IngredientCategoryAlreadyExists.Error(), http.StatusConflict)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.IngredientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *IngredientUseCase) DeleteIngredientCategory(id uint) *customErr.CustomError {
	if err := u.repoIngredient.DeleteIngredientCategory(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.IngredientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *IngredientUseCase) CreateIngredient(ingredient *models.Ingredient) (uint, *customErr.CustomError) {
	if _, err := u.repoIngredient.GetIngredientCategoryByID(ingredient.IngredientCategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, customErr.NewCustomError(err, customErr.IngredientCategoryNotFound.Error(), http.StatusNotFound)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	id, err := u.repoIngredient.CreateIngredient(ingredient)
	if err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return 0, customErr.NewCustomError(err, customErr.IngredientAlreadyExists.Error(), http.StatusConflict)
		} else {
			return 0, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return id, nil
}

func (u *IngredientUseCase) GetAllIngredients() (*[]models.Ingredient, *customErr.CustomError) {
	ingredients, err := u.repoIngredient.GetAllIngredients()
	if err != nil {
		return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
	}

	return ingredients, nil
}

func (u *IngredientUseCase) GetIngredientByID(id uint) (*models.Ingredient, *customErr.CustomError) {
	ingredient, err := u.repoIngredient.GetIngredientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
		} else {
			return nil, customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return ingredient, nil
}

func (u *IngredientUseCase) UpdateIngredient(ingredient *models.Ingredient) *customErr.CustomError {

	if ingredient.IngredientCategoryID != 0 {
		if _, err := u.repoIngredient.GetIngredientCategoryByID(ingredient.IngredientCategoryID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return customErr.NewCustomError(err, customErr.IngredientCategoryNotFound.Error(), http.StatusNotFound)
			} else {
				return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
			}
		}
	}

	ingredient.UpdatedAt = time.Now()

	if err := u.repoIngredient.UpdateIngredient(ingredient); err != nil {
		if ok, _ := customErr.IsDuplicateKeyError(err); ok {
			return customErr.NewCustomError(err, customErr.IngredientAlreadyExists.Error(), http.StatusConflict)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (u *IngredientUseCase) DeleteIngredient(id uint) *customErr.CustomError {
	if err := u.repoIngredient.DeleteIngredient(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customErr.NewCustomError(err, customErr.IngredientNotFound.Error(), http.StatusNotFound)
		} else {
			return customErr.NewCustomError(err, customErr.ServerError.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}
