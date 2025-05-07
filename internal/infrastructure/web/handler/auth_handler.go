package handler

import (
	"github.com/ParthSharma272/GoStock/internal/domain/common"                 // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/middleware" // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/request"    // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/response"   // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/service"                       // Replace your_project_path
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	// Default role to customer if not provided or invalid
	role := req.Role
	if role != common.RoleAdmin && role != common.RoleCustomer && role != common.RoleShipper {
		role = common.RoleCustomer
	}
	// For security, usually admin registration should be a separate, protected endpoint or manual process.
	// Here, we'll allow it for simplicity but restrict actual admin actions via middleware.
	if role == common.RoleAdmin {
		// Potentially add a check here, e.g., a secret registration key for the first admin
	}

	user, token, err := h.authService.Register(req.FirstName, req.LastName, req.Email, req.Password, role)
	if err != nil {
		c.JSON(http.StatusConflict, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"token":   token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"token":   token,
	})
}

func (h *AuthHandler) GetMyProfile(c *gin.Context) {
	claims := middleware.GetAuthClaims(c)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: "User not authenticated"})
		return
	}

	user, err := h.authService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.StandardErrorResponse{Error: "User not found"})
		return
	}

	// Don't send password hash in profile
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}
