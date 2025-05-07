package service

import (
	"errors"
	domainProduct "github.com/ParthSharma272/GoStock/internal/domain/product" // Replace your_project_path
	"gorm.io/gorm"
)

type productService struct {
	db   *gorm.DB
	repo domainProduct.Repository
	// categoryRepo category.Repository // Inject if needed for validation
}

func NewProductService(db *gorm.DB, repo domainProduct.Repository) domainProduct.Service {
	return &productService{db: db, repo: repo}
}

func (s *productService) CreateProduct(name, description string, price float64, stock uint, categoryID uint) (*domainProduct.Product, error) {
	if name == "" {
		return nil, errors.New("product name is required")
	}
	if price <= 0 {
		return nil, errors.New("product price must be positive")
	}
	// Add category existence check if categoryRepo is injected

	product := &domainProduct.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CategoryID:  categoryID,
	}
	err := s.repo.Create(s.db, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetProductByID(id uint) (*domainProduct.Product, error) {
	return s.repo.FindByID(s.db, id)
}

func (s *productService) GetAllProducts(page, pageSize int) ([]domainProduct.Product, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 { // Max page size
		pageSize = 100
	}
	offset := (page - 1) * pageSize
	return s.repo.FindAll(s.db, offset, pageSize)
}

func (s *productService) UpdateProduct(id uint, name *string, description *string, price *float64, stock *uint, categoryID *uint) (*domainProduct.Product, error) {
	product, err := s.repo.FindByID(s.db, id)
	if err != nil {
		return nil, err // Could be gorm.ErrRecordNotFound
	}

	if name != nil && *name != "" {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if price != nil && *price > 0 {
		product.Price = *price
	}
	if stock != nil { // Allows setting stock to 0
		product.Stock = *stock
	}
	if categoryID != nil && *categoryID > 0 {
		product.CategoryID = *categoryID
	}

	err = s.repo.Update(s.db, product)
	return product, err
}

func (s *productService) DeleteProduct(id uint) error {
	// Add logic here if you need to check for orders associated with the product, etc.
	return s.repo.Delete(s.db, id)
}
