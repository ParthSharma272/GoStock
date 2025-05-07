package user

import (
	"github.com/ParthSharma272/GoStock/internal/domain/common" // Replace your_project_path

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string          `json:"first_name" gorm:"not null"`
	LastName  string          `json:"last_name" gorm:"not null"`
	Email     string          `json:"email" gorm:"uniqueIndex;not null"`
	Password  string          `json:"-" gorm:"not null"`
	Role      common.UserRole `json:"role" gorm:"type:varchar(20);not null;default:'customer'"`
	IsActive  bool            `json:"is_active" gorm:"default:true"`
	// Orders    []order.Order `json:"-" gorm:"foreignKey:CustomerID"` // Add when Order domain is complete
}

// BeforeSave GORM hook to hash password
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" && len(u.Password) < 60 { // Avoid re-hashing an already hashed password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
