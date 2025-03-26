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

// login contorller interface
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
	var userfromout struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ctx.ShouldBind(&userfromout)
	var user entity.User
	if err := controller.db.Where("email = ?", userfromout.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	isUserAuthenticated := controller.loginService.LoginUser(user.Email, user.Password)
	if !isUserAuthenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	}
	accessToken := controller.jWtService.GenerateToken(user.Email, user.Id, time.Minute*10)
	refreshToken := controller.jWtService.GenerateToken(user.Email, user.Id, time.Hour*1)
	controller.db.Create(&entity.UserToken{
		UserID:       user.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Minute * 10),
	})
	ctx.JSON(http.StatusOK, gin.H{
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
	// Tạo access token mới với thông tin từ refresh token
	newAccessToken := controller.jWtService.GenerateToken(claims.Email, claims.UserID, time.Minute*10)
	controller.db.Model(&entity.UserToken{}).
		Where("refresh_token = ?", refreshToken).
		Update("access_token", newAccessToken)
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}
