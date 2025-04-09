package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
	"vietha/src/entity"
	"vietha/src/middleware"
	"vietha/src/service"
	"vietha/src/utils"
)

type Auths interface {
	ForgotPassword(c *gin.Context)
	ProfileHandler(c *gin.Context)
	RegisterAuthRoutes(router *gin.Engine)
}
type Auth struct {
	db           *gorm.DB
	authorizeJWT middleware.JWTMiddleware
	service      service.JWTService
}

func NewAuth(db *gorm.DB, authorizeJWT middleware.JWTMiddleware, jwtService service.JWTService) *Auth {
	return &Auth{
		db:           db,
		service:      jwtService,
		authorizeJWT: authorizeJWT,
	}
}

var tokenStore = sync.Map{}

func GenerateToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return hex.EncodeToString(bytes)
}
func (auth *Auth) ForgotPassword(c *gin.Context) {
	var emailjson struct {
		Email string `json:"email"`
	}
	c.ShouldBind(&emailjson)
	email := emailjson.Email
	var user entity.User
	result := auth.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email don't exist"})
		return
	}
	token := GenerateToken()
	tokenStore.Store(token, email)

	go func() {
		time.Sleep(15 * time.Minute)
		tokenStore.Delete(token)
	}()
	err := utils.SendResetPasswordEmail(email, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not send mail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email has been sent"})
}

type jwtController struct {
	db           *gorm.DB
	authorizeJWT middleware.JWTMiddleware
	service      service.JWTService
}

// Constructor

// Method trả thông tin người dùng từ token
func (ctrl *jwtController) ProfileHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	authClaims, ok := claims.(*service.AuthCustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "This is validate token",
		"email":       authClaims.Email,
		"name":        authClaims.Name,
		"userId":      authClaims.UserID,
		"issuer":      authClaims.Issuer,
		"expires":     authClaims.ExpiresAt,
		"roles":       authClaims.Roles,
		"permissions": authClaims.Permissions,
	})
}

// Đăng ký route có bảo vệ bằng JWT
