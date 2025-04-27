package auth

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type LogoutController interface {
	Logout(ctx *gin.Context)
	LogoutAllDevices(ctx *gin.Context)
}
type logoutController struct {
	tokenService service.TokenService
}
type LogoutAllDevicesRequest struct {
	UserID int `json:"user_id"`
}

func NewLogout(tokenService service.TokenService) LogoutController {
	return &logoutController{
		tokenService: tokenService,
	}
}

// Logout @Summary Đăng xuất khỏi thiết bị hiện tại
// @Description Xóa access token hiện tại (thiết bị đang dùng)
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} map[string]interface{} "Đăng xuất thành công"
// @Failure 40001 {object} map[string]interface{} "Lỗi khi xóa token"
// @Router /api/v1/auth/logout [post]
func (controller *logoutController) Logout(ctx *gin.Context) {
	// Nếu logout trên thiết bị thì lưu token vào redis
	token := ctx.GetHeader("Authorization")
	if token == "" {
		global.Logger.Warn("Authorization header rỗng")
		response.ErrorResponse(ctx, 20002)
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	err := controller.tokenService.DeleteAccessToken(token)

	if err != nil {
		global.Logger.Warn("Lỗi khi xóa token", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}
	global.Logger.Info("Đăng xuất thành công thiết bị")
	response.SuccessResponse(ctx, 20001, "Đăng xuất khỏi thiết bị thành công ")
	return
}

// LogoutAllDevices @Summary Đăng xuất khỏi tất cả thiết bị
// @Description Xóa tất cả access token của người dùng dựa vào user_id
// @Tags Auth
// @Accept json
// @Produce json
// @Param user_id body LogoutAllDevicesRequest true "ID người dùng cần đăng xuất toàn bộ"
// @Success 200 {object} map[string]interface{} "Đăng xuất thành công"
// @Failure 20003 {object} map[string]interface{} "User ID không hợp lệ"
// @Failure 40001 {object} map[string]interface{} "Lỗi khi xóa tất cả token"
// @Router /api/v1/auth/logout-all [post]
func (controller *logoutController) LogoutAllDevices(ctx *gin.Context) {
	// Nếu logout hết tất cả thiết bị thì xóa hết Access Token có trong thiết bị đó
	var request LogoutAllDevicesRequest
	if err := ctx.ShouldBindJSON(&request); err != nil || request.UserID == 0 {
		global.Logger.Warn("User Id không hợp lệ")
		response.ErrorResponse(ctx, 20003)
		return
	}
	err := controller.tokenService.DeleteAccessTokenByUserId(request.UserID)
	if err != nil {
		global.Logger.Error("Lỗi khi xóa tất cả token", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}
	global.Logger.Info("Đăng xuất khỏi tất cả thiết bị thành công", zap.Int("user_id", request.UserID))
	response.SuccessResponse(ctx, 20001, "Đăng xuất khỏi tất cả thiết bị thành công")
	return
}
