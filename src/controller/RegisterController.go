package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vietha/src/service"
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
func (controller *registerController) Register(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}
	user, err := controller.userService.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Dữ liệu đầu vào:", input)
	ctx.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công", "user": user})
}
