package user

import "gorm.io/gorm"

type Repository interface {
	Create(db *gorm.DB, user *User) error
	FindByEmail(db *gorm.DB, email string) (*User, error)
	FindByID(db *gorm.DB, id uint) (*User, error)
	// Add other methods like Update, Delete, GetAll as needed for admin user management
}
