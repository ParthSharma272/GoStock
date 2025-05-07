package product

import (
	// "your_project_path/internal/domain/category" // Add when Category domain is complete
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null;type:decimal(10,2)"`
	Stock       uint    `json:"stock" gorm:"not null;default:0"`
	CategoryID  uint    `json:"category_id"`
	// Category    category.Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"` // Belongs to Category
}
