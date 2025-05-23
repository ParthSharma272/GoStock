package product

import "gorm.io/gorm"

type Repository interface {
	Create(db *gorm.DB, product *Product) error
	FindByID(db *gorm.DB, id uint) (*Product, error)
	FindAll(db *gorm.DB, offset, limit int) ([]Product, int64, error)
	Update(db *gorm.DB, product *Product) error
	Delete(db *gorm.DB, id uint) error
	// UpdateStock(tx *gorm.DB, productID uint, quantityChange int) error
}
