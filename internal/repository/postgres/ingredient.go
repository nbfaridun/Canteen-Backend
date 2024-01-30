package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type IngredientPostgres struct {
	db *gorm.DB
}

func NewIngredientPostgres(db *gorm.DB) *IngredientPostgres {
	return &IngredientPostgres{db: db}
}

func (r *IngredientPostgres) CreateIngredientCategory(ingredientCategory *models.IngredientCategory) (uint, error) {
	result := r.db.Table(constants.IngredientCategoryTableName).Create(ingredientCategory)
	if result.Error != nil {
		return 0, result.Error
	}

	return ingredientCategory.ID, nil
}

func (r *IngredientPostgres) GetAllIngredientCategories() (*[]models.IngredientCategory, error) {
	var ingredientCategories []models.IngredientCategory
	result := r.db.Table(constants.IngredientCategoryTableName).Find(&ingredientCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ingredientCategories, nil
}

func (r *IngredientPostgres) GetIngredientCategoryByID(id uint) (*models.IngredientCategory, error) {
	var ingredientCategory models.IngredientCategory
	result := r.db.Table(constants.IngredientCategoryTableName).Where("ingredient_category_id = ?", id).First(&ingredientCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ingredientCategory, nil
}

func (r *IngredientPostgres) UpdateIngredientCategory(ingredientCategory *models.IngredientCategory) error {
	result := r.db.Table(constants.IngredientCategoryTableName).Where("ingredient_category_id = ?", ingredientCategory.ID).Updates(ingredientCategory)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *IngredientPostgres) DeleteIngredientCategory(id uint) error {
	result := r.db.Table(constants.IngredientCategoryTableName).Where("ingredient_category_id = ?", id).Delete(&models.IngredientCategory{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *IngredientPostgres) CreateIngredient(ingredient *models.Ingredient) (uint, error) {
	result := r.db.Table(constants.IngredientTableName).Create(ingredient)
	if result.Error != nil {
		return 0, result.Error
	}

	return ingredient.ID, nil
}

func (r *IngredientPostgres) GetAllIngredients() (*[]models.Ingredient, error) {
	var ingredients []models.Ingredient
	result := r.db.Table(constants.IngredientTableName).Find(&ingredients)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ingredients, nil
}

func (r *IngredientPostgres) GetIngredientByID(id uint) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	result := r.db.Table(constants.IngredientTableName).Where("ingredient_id = ?", id).First(&ingredient)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ingredient, nil
}

func (r *IngredientPostgres) UpdateIngredient(ingredient *models.Ingredient) error {
	result := r.db.Table(constants.IngredientTableName).Where("ingredient_id = ?", ingredient.ID).Updates(ingredient)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *IngredientPostgres) DeleteIngredient(id uint) error {
	result := r.db.Table(constants.IngredientTableName).Delete(&models.Ingredient{}, "ingredient_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
