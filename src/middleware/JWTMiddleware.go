package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vietha/src/service"
)

// AuthorizeJWT - Middleware kiểm tra JWT
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
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
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// Lưu claims vào context với key "claims"
		c.Set("claims", claims)
		c.Next()
	}
}
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing claims"})
			return
		}
		authClaims, ok := claims.(*service.AuthCustomClaims)
		var roles []string
		for i, role := range authClaims.Roles {
			roles[i] = role.Name
		}
		for _, roleOut := range allowedRoles {
			for _, roleIn := range roles {
				if roleOut == roleIn {
					c.Next()
					return
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Don't have role to perform this action"})
	}
}
func PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing claims"})
			return
		}
		authClaims, ok := claims.(*service.AuthCustomClaims)
		var permissions []string
		for _, permission := range authClaims.Permissions {
			permissions = append(permissions, permission.Name)
		}
		for _, permission := range allowedPermissions {
			for _, permissionIn := range permissions {
				if permissionIn == permission {
					c.Next()
					return
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Don't have permission to perform this action"})
	}
}
