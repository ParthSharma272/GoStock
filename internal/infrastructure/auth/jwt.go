package auth

import (
	"fmt"
	"github.com/ParthSharma272/GoStock/internal/domain/common"         // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/domain/user"           // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config" // Replace your_project_path
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint            `json:"user_id"`
	Email  string          `json:"email"`
	Role   common.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(appUser *user.User, cfg *config.Config) (string, error) {
	expirationTime := time.Now().Add(cfg.JWTExpiration)
	claims := &Claims{
		UserID: appUser.ID,
		Email:  appUser.Email,
		Role:   common.UserRole(appUser.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-ecommerce-backend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecretKey))
}

func ValidateToken(tokenString string, cfg *config.Config) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
