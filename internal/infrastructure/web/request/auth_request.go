package request

import "github.com/ParthSharma272/GoStock/internal/domain/common" // Replace your_project_path

type RegisterRequest struct {
	FirstName string          `json:"first_name" binding:"required"`
	LastName  string          `json:"last_name" binding:"required"`
	Email     string          `json:"email" binding:"required,email"`
	Password  string          `json:"password" binding:"required,min=6"`
	Role      common.UserRole `json:"role"` // Optional, defaults in DB or service
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
