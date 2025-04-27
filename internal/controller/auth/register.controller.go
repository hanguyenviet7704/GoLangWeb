package auth

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterController interface {
	Register(ctx *gin.Context)
}
type registerController struct {
	userService service.UserService
}

func NewRegisterController(userService service.UserService) RegisterController {
	return &registerController{
		userService: userService,
	}
}

// Register @Summary Đăng ký người dùng mới
// @Description Đăng ký tài khoản mới với thông tin tên, email, và mật khẩu.
// @Tags Auth
// @Accept json
// @Produce json
// @Param name body string true "Tên người dùng"
// @Param email body string true "Email người dùng"
// @Param password body string true "Mật khẩu người dùng"
// @Success 200 {object} map[string]interface{} "Đăng ký thành công"
// @Failure 20002 {object} map[string]interface{} "Dữ liệu không hợp lệ"
// @Failure 50002 {object} map[string]interface{} "Email đã tồn tại"
// @Failure 40001 {object} map[string]interface{} "Lỗi khi đăng ký"
// @Router /api/v1/auth/register [post]

func (controller *registerController) Register(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		global.Logger.Warn("Dữ liệu không hợp lệ", zap.Error(err))
		response.ErrorResponse(ctx, 20002)
		return
	}
	user, err := controller.userService.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {

		switch err.Error() {
		case "EMAIL_EXIST":
			global.Logger.Error("Lỗi khi đăng ký: %v", zap.Error(err))
			response.ErrorResponse(ctx, 50002)
			return
		default:
			global.Logger.Error("Lỗi khi đăng kí user với lỗi từ database")
			response.ErrorResponse(ctx, 40001)
			return
		}
	}
	global.Logger.Info("Đăng ký thành công cho ", zap.String("email", user.Email))
	response.SuccessResponse(ctx, 20001, nil)
}
