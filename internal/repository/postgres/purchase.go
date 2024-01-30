package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type PurchasePostgres struct {
	db *gorm.DB
}

func NewPurchasePostgres(db *gorm.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r *PurchasePostgres) CreateSupplier(supplier *models.Supplier) (uint, error) {
	result := r.db.Table(constants.SupplierTableName).Create(supplier)
	if result.Error != nil {
		return 0, result.Error
	}

	return supplier.ID, nil
}

func (r *PurchasePostgres) GetAllSuppliers() (*[]models.Supplier, error) {
	var suppliers []models.Supplier
	result := r.db.Table(constants.SupplierTableName).Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return &suppliers, nil
}

func (r *PurchasePostgres) GetSupplierByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	result := r.db.Table(constants.SupplierTableName).Where("supplier_id = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}

func (r *PurchasePostgres) UpdateSupplier(supplier *models.Supplier) error {
	result := r.db.Table(constants.SupplierTableName).Where("supplier_id = ?", supplier.ID).Updates(supplier)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PurchasePostgres) DeleteSupplier(id uint) error {
	result := r.db.Table(constants.SupplierTableName).Where("supplier_id = ?", id).Delete(&models.Supplier{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PurchasePostgres) CreatePurchase(purchase *models.Purchase) (uint, error) {
	result := r.db.Table(constants.PurchaseTableName).Create(purchase)
	if result.Error != nil {
		return 0, result.Error
	}

	return purchase.ID, nil
}

func (r *PurchasePostgres) CreatePurchasesIngredients(purchasesIngredients *[]models.PurchasesIngredients) error {
	result := r.db.Table(constants.PurchasesIngredientsTableName).Create(purchasesIngredients)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
