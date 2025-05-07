package postgres

import (
	"errors"
	"github.com/ParthSharma272/GoStock/internal/domain/product" // Replace your_project_path

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct{}

func NewProductRepository() product.Repository {
	return &productRepository{}
}

func (r *productRepository) Create(db *gorm.DB, p *product.Product) error {
	return db.Create(p).Error
}

func (r *productRepository) FindByID(db *gorm.DB, id uint) (*product.Product, error) {
	var p product.Product
	// err := db.Preload("Category").First(&p, id).Error // Uncomment when Category is ready
	err := db.First(&p, id).Error
	return &p, err
}

func (r *productRepository) FindAll(db *gorm.DB, offset, limit int) ([]product.Product, int64, error) {
	var products []product.Product
	var total int64

	db.Model(&product.Product{}).Count(&total)
	// err := db.Preload("Category").Offset(offset).Limit(limit).Order("created_at desc").Find(&products).Error // Uncomment Category
	err := db.Offset(offset).Limit(limit).Order("created_at desc").Find(&products).Error
	return products, total, err
}

func (r *productRepository) Update(db *gorm.DB, p *product.Product) error {
	return db.Save(p).Error
}

func (r *productRepository) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&product.Product{}, id).Error
}

func (r *productRepository) UpdateStock(tx *gorm.DB, productID uint, quantityChange int) error {
	var p product.Product
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&p, productID).Error; err != nil {
		return err // Product not found or other DB error
	}

	newStock := int(p.Stock) + quantityChange
	if newStock < 0 {
		return errors.New("insufficient stock")
	}
	p.Stock = uint(newStock)
	return tx.Save(&p).Error
}
