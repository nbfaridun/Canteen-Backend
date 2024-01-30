package response

import "Canteen-Backend/internal/models"

type GetSupplier struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func MapSupplierToGetSupplier(supplier *models.Supplier) *GetSupplier {
	return &GetSupplier{
		ID:   supplier.ID,
		Name: supplier.Name,
	}
}
