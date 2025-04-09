package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vietha/src/service"
)

type JWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
	RoleMiddleware(allowedRoles ...string) gin.HandlerFunc
	PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc
}
type middleware struct {
	jwtService service.JWTService
}

func MiddlewareHandler(jwtService service.JWTService) JWTMiddleware {
	return &middleware{
		jwtService: jwtService,
	}
}

// AuthorizeJWT - Middleware kiểm tra JWT
func (m *middleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			return
		}
		const BEARER_SCHEMA = "Bearer "
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// RoleMiddleware - Middleware kiểm tra vai trò
func (m *middleware) RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing claims"})
			return
		}

		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		roles := make([]string, 0, len(authClaims.Roles))
		for _, role := range authClaims.Roles {
			roles = append(roles, role.Name)
		}

		for _, allowed := range allowedRoles {
			for _, actual := range roles {
				if allowed == actual {
					c.Next()
					return
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have required role"})
	}
}

// PermissionMiddleware - Middleware kiểm tra quyền
func (m *middleware) PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing claims"})
			return
		}

		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		permissions := make([]string, 0, len(authClaims.Permissions))
		for _, perm := range authClaims.Permissions {
			permissions = append(permissions, perm.Name)
		}

		for _, allowed := range allowedPermissions {
			for _, actual := range permissions {
				if allowed == actual {
					c.Next()
					return
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have required permission"})
	}
}
