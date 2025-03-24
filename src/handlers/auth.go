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
	"vietha/src/utils"
)

type Auths interface {
	ForgotPassword(c *gin.Context)
}
type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db: db}
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
