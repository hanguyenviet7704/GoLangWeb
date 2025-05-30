package auth

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

// http://localhost:8080/api/auth/register
type LoginController interface {
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}
type loginController struct {
	tokenService service.TokenService
	userService  service.UserService
	jWtService   service.JWTService
}

func NewLoginController(tokenService service.TokenService, userService service.UserService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		tokenService: tokenService,
		userService:  userService,
		jWtService:   jWtService,
	}
}

// Login @Summary Đăng nhập người dùng
// @Description Đăng nhập vào hệ thống với email và mật khẩu để nhận Access Token và Refresh Token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "Email người dùng"
// @Param password body string true "Mật khẩu người dùng"
// @Success 200 {object} map[string]interface{} "Đăng nhập thành công"
// @Failure 50001 {object} map[string]interface{} "Không tìm thấy người dùng"
// @Failure 50003 {object} map[string]interface{} "Sai mật khẩu"
// @Failure 40004 {object} map[string]interface{} "Lỗi tạo token"
// @Failure 40001 {object} map[string]interface{} "Lỗi hệ thống"
// @Router /api/v1/auth/login [post]
func (controller *loginController) Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ctx.ShouldBind(&input)
	user, accessToken, refreshToken, err := controller.userService.LoginUser(ctx, input.Email, input.Password)
	if err != nil {
		switch err.Error() {
		case "USER_NOT_FOUND":
			global.Logger.Warn("Không tìm thấy người dùng")
			response.ErrorResponse(ctx, 50001)
		case "WRONG_PASSWORD":
			global.Logger.Warn("Sai mật khẩu của tài khoản")
			response.ErrorResponse(ctx, 50003)
		case "ACCESS_TOKEN_ERROR", "REFRESH_TOKEN_ERROR":
			global.Logger.Error("Lỗi khi tạo Token")
			response.ErrorResponse(ctx, 40004)
		default:
			global.Logger.Error("Lỗi hệ thống khi Login")
			response.ErrorResponse(ctx, 40001)
		}
		return
	}

	data := gin.H{
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
			"name":  user.Name,
		},
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	global.Logger.Info("Đăng nhập thành công", zap.Any("data", data))
	response.SuccessResponse(ctx, 20001, data)
	return
}

// RefreshToken @Summary Làm mới Access Token
// @Description Sử dụng refreshToken để lấy accessToken mới.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Refresh Token"
// @Success 200 {object} map[string]interface{} "Token mới"
// @Failure 30002 {object} map[string]interface{} "Token không hợp lệ"
// @Failure 40002 {object} map[string]interface{} "Lỗi khi truy vấn Refresh Token"
// @Failure 40003 {object} map[string]interface{} "Không thể tạo Access Token"
// @Router /api/v1/auth/refresh [post]
func (controller *loginController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Authorization")
	if refreshToken == "" {
		global.Logger.Warn("Token không hợp lệ")
		response.ErrorResponse(ctx, 30002)
		return
	}
	refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")
	claims, err := controller.jWtService.ValidateToken(refreshToken)
	if err != nil {
		global.Logger.Error("Token hết hạn")
		response.ErrorResponse(ctx, 30002)
		return
	}
	err = controller.tokenService.FindRefreshToken(refreshToken)
	if err != nil {
		global.Logger.Error("Lỗi khi truy vấn RefreshToken")
		response.ErrorResponse(ctx, 40002)
		return
	}
	var user *po.User
	user, err = controller.userService.GetUserById(claims.UserID)
	if err != nil {
		global.Logger.Error("Không tìm thấy người dùng ", zap.String("email", claims.Email))
		response.ErrorResponse(ctx, 40002)
		return
	}
	// Tạo access token mới với thông tin từ refresh token
	newAccessToken, err := controller.jWtService.GenerateAccessToken(user)
	if err != nil {
		global.Logger.Error("Không thể tạo AccessToken")
		response.ErrorResponse(ctx, 40003)
		return
	}
	err = controller.tokenService.UpdateAccessToken(refreshToken, newAccessToken)
	if err != nil {
		global.Logger.Error("Lỗi khi truy vấn để update Access Token")
		response.ErrorResponse(ctx, 40002)
		return
	}
	global.Logger.Info("Refresh Token thành công ", zap.Any("refreshToken", refreshToken))
	response.SuccessResponse(ctx, 20001, gin.H{
		"accessToken": newAccessToken,
	})
}
