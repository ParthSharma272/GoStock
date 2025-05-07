package middleware

import (
	"github.com/ParthSharma272/GoStock/internal/domain/common"                 // Replace your_project_path
	infraAuth "github.com/ParthSharma272/GoStock/internal/infrastructure/auth" // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config"         // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/response"   // Replace your_project_path
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeaderKey)
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: "Authorization header is not provided"})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: "Invalid authorization header format"})
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != AuthorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: "Unsupported authorization type: " + authType})
			return
		}

		accessToken := fields[1]
		claims, err := infraAuth.ValidateToken(accessToken, cfg)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.StandardErrorResponse{Error: "Invalid or expired token: " + err.Error()})
			return
		}

		c.Set(AuthorizationPayloadKey, claims)
		c.Next()
	}
}

func GetAuthClaims(c *gin.Context) *infraAuth.Claims {
	val, exists := c.Get(AuthorizationPayloadKey)
	if !exists {
		return nil
	}
	claims, ok := val.(*infraAuth.Claims)
	if !ok {
		return nil
	}
	return claims
}

func RoleMiddleware(requiredRoles ...common.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := GetAuthClaims(c)
		if claims == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, response.StandardErrorResponse{Error: "Access denied: User claims not found."})
			return
		}

		authorized := false
		for _, role := range requiredRoles {
			if claims.Role == role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, response.StandardErrorResponse{Error: "Access denied: You do not have the required role."})
			return
		}
		c.Next()
	}
}
