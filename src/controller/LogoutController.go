package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"vietha/src/entity"
)

type LogoutController interface {
	Logout(ctx *gin.Context)
	LogoutAllDevices(ctx *gin.Context)
}
type logoutController struct {
	db *gorm.DB
}

func NewLogout(db *gorm.DB) LogoutController {
	return &logoutController{
		db: db,
	}
}
func (controller *logoutController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is empty"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	result := controller.db.Where("access_token = ?", token).Delete(&entity.UserToken{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out from this device"})
}

func (controller *logoutController) LogoutAllDevices(ctx *gin.Context) {
	var request struct {
		UserID int `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil || request.UserID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	result := controller.db.Where("user_id = ?", request.UserID).Delete(&entity.UserToken{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out from all devices"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out from all devices"})
}
