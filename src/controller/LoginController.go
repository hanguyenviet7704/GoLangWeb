package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
	"vietha/src/entity"
	"vietha/src/service"
)

type LoginController interface {
	Login(ctx *gin.Context) // trả về JWT nếu đăng nhập thành công hoạc một thông báo lỗi nếu thất bại
	RefreshToken(ctx *gin.Context)
}
type loginController struct {
	db           *gorm.DB
	loginService service.UserService
	jWtService   service.JWTService
}

func LoginHandler(db *gorm.DB, loginService service.UserService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		db:           db,
		loginService: loginService,
		jWtService:   jWtService,
	}
}
func (controller *loginController) Login(ctx *gin.Context) {
	var userRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ctx.ShouldBind(&userRequest)
	var user entity.User
	if err := controller.db.Where("email = ?", userRequest.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	isUserAuthenticated := controller.loginService.LoginUser(user.Email, user.Password)
	if !isUserAuthenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	}
	accessToken := controller.jWtService.GenerateAccessToken(&user)
	refreshToken := controller.jWtService.GenerateRefreshToken(&user)
	controller.db.Create(&entity.Tokens{
		UserID:                   user.Id,
		AccessToken:              accessToken,
		RefreshToken:             refreshToken,
		Access_token_expires_at:  time.Now().Add(time.Minute * 10),
		Refresh_token_expires_at: time.Now().Add(time.Hour * 1),
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	})
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
			"name":  user.Name,
		},
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
func (controller *loginController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Authorization")
	if refreshToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is missing"})
		return
	}
	refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")
	claims, err := controller.jWtService.ValidateToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid or expired"})
		return
	}
	err = controller.db.Where("refresh_token = ?", refreshToken).First(&refreshToken).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid in Database "})
		return
	}
	var user entity.User
	controller.db.Where("claims.user_id = ?", claims.UserID).First(&user)
	// Tạo access token mới với thông tin từ refresh token
	newAccessToken := controller.jWtService.GenerateAccessToken(&user)
	controller.db.Model(&entity.Tokens{}).
		Where("refresh_token = ?", refreshToken).
		Update("access_token", newAccessToken)
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}
