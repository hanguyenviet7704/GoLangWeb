package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"vietha/src/entity"
)

type LogoutController interface {
	Logout(ctx *gin.Context)
	LogoutAllDevices(ctx *gin.Context)
	LogoutDevice(ctx *gin.Context)
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
	result := controller.db.Where("access_token = ?", token).Delete(&entity.UserToken{})
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is empty"})
		return
	}
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success log out in now device"})
}
func (controller *logoutController) LogoutAllDevices(ctx *gin.Context) {
	userID := make(map[string]int)
	if err := ctx.ShouldBindJSON(&userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// Xóa tất cả token của user
	result := controller.db.Where("user_id = ?", userID["user_id"]).Delete(&entity.UserToken{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out from all devices"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out from all devices"})
}
func (controller *logoutController) LogoutDevice(ctx *gin.Context) {
	var req struct {
		UserID   int    `form:"user_id"`
		DeviceID string `json:"device_id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Xóa token của thiết bị cụ thể
	result := controller.db.Where("user_id = ? AND device_id = ?", req.UserID, req.DeviceID).Delete(&entity.UserToken{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out from device"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out from device " + req.DeviceID})
}
