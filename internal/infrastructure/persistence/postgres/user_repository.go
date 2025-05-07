package postgres

import (
	"github.com/ParthSharma272/GoStock/internal/domain/user" // Replace your_project_path
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (r *userRepository) Create(db *gorm.DB, u *user.User) error {
	return db.Create(u).Error
}

func (r *userRepository) FindByEmail(db *gorm.DB, email string) (*user.User, error) {
	var u user.User
	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) FindByID(db *gorm.DB, id uint) (*user.User, error) {
	var u user.User
	if err := db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
