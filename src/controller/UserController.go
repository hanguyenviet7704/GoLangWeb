package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vietha/src/entity"
	"vietha/src/service"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateOrUpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetRolesFromUser(ctx *gin.Context)
	AssignRolesToUser(c *gin.Context)
}
type userController struct {
	userservice service.UserService
}

func NewUserController(userservice service.UserService) UserController {
	return &userController{
		userservice: userservice,
	}
}

// GetAllUsers godoc
// @Summary Lấy danh sách người dùng
// @Description Trả về danh sách tất cả người dùng trong hệ thống
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]entity.User
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func (controller *userController) GetAllUsers(c *gin.Context) {
	users, err := controller.userservice.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if users == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Don't have any users let creat "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (controller *userController) GetUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when converting string to int"})
	}
	user, err := controller.userservice.GetUserById(idNum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (controller *userController) CreateOrUpdateUser(c *gin.Context) {
	var input struct {
		ID       string   `json:"id" `
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		Password string   `json:"password"`
		IsActive bool     `json:"isActive"`
		Role     []string `json:"role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bind JSON"})
		return
	}
	idNum, err := strconv.Atoi(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when converting string to int"})
	}
	user := &entity.User{
		Id:        idNum,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Is_active: input.IsActive,
		UpdatedAt: time.Now(),
	}
	err = controller.userservice.CreateOrUpdateUser(user, input.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Create Or Update User Successful": user})
	return
}
func (controller *userController) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when converting string to int"})
	}
	err = controller.userservice.DeleteUser(idNum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Delete User Successful": id})
}
func (controller *userController) GetRolesFromUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when converting string to int"})
		return
	}
	roles, err := controller.userservice.GetRoleFromUser(idNum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if roles == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User do not have any roles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  id,
		"roles": roles,
	})
	return
}
func (controller *userController) AssignRolesToUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when converting string to int"})
		return
	}
	var input struct {
		RoleNames []string `json:"roleNames"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}
	err = controller.userservice.AssignRolesToUser(idNum, input.RoleNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Assign Role Successful", "userId": idNum})
}
