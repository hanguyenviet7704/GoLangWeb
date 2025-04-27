package controller

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type RolesController interface {
	GetAllRoles(ctx *gin.Context)
	GetRoleById(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	GetPermissionsFromRole(ctx *gin.Context)
	AssignPermissionsToRole(ctx *gin.Context)
}
type rolesController struct {
	roleService service.RoleService
}

func NewRolesController(roleService service.RoleService) RolesController {
	return &rolesController{
		roleService: roleService,
	}
}

// GetAllRoles @Summary Lấy tất cả các Role
// @Description Trả về danh sách tất cả các Role trong hệ thống.
// @Tags Roles
// @Produce json
// @Success 200 {object} map[string]interface{} "Lấy danh sách role thành công"
// @Failure 40001 {object} map[string]interface{} "Lỗi khi lấy danh sách roles"
// @Router /api/v1/roles [get]
func (c *rolesController) GetAllRoles(ctx *gin.Context) {
	roles, err := c.roleService.GetAllRoles()
	if err != nil {
		global.Logger.Error("Lỗi khi lấy danh sách roles: ", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"roles": roles})
}

// GetRoleById @Summary Lấy Role theo ID
// @Description Lấy thông tin chi tiết của Role theo ID.
// @Tags Roles
// @Produce json
// @Param id path int true "ID Role"
// @Success 200 {object} map[string]interface{} "Lấy role theo ID thành công"
// @Failure 40001 {object} map[string]interface{} "ID không hợp lệ"
// @Failure 40002 {object} map[string]interface{} "Lỗi khi lấy role theo ID"
// @Router /api/v1/roles/{id} [get]
func (c *rolesController) GetRoleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	role, err := c.roleService.GetRoleByID(id)
	if err != nil {
		global.Logger.Error("Lỗi khi lấy role theo ID: ", zap.Error(err))
		response.ErrorResponse(ctx, 40002)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"role": role})
}

// CreateRole @Summary Tạo mới Role
// @Description Thêm một Role mới vào hệ thống.
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body po.Roles true "Thông tin Role cần tạo"
// @Success 201 {object} map[string]interface{} "Tạo role thành công"
// @Failure 40003 {object} map[string]interface{} "Lỗi bind JSON"
// @Failure 40004 {object} map[string]interface{} "Lỗi khi tạo role"
// @Router /api/v1/roles [post]
func (c *rolesController) CreateRole(ctx *gin.Context) {
	var role po.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		global.Logger.Error("Lỗi khi bind JSON tạo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	if err := c.roleService.CreateRole(&role); err != nil {
		global.Logger.Error("Lỗi khi tạo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40004)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"role": role})
}

// UpdateRole @Summary Cập nhật Role
// @Description Cập nhật thông tin Role theo ID.
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "ID Role cần cập nhật"
// @Param role body po.Roles true "Thông tin Role cập nhật"
// @Success 200 {object} map[string]interface{} "Cập nhật role thành công"
// @Failure 40001 {object} map[string]interface{} "ID không hợp lệ"
// @Failure 40003 {object} map[string]interface{} "Lỗi bind JSON"
// @Failure 40004 {object} map[string]interface{} "Lỗi khi cập nhật role"
// @Router /api/v1/roles/{id} [put]
func (c *rolesController) UpdateRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	var role po.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		global.Logger.Error("Lỗi khi bind JSON cập nhật role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	role.ID = id
	if err := c.roleService.UpdateRole(&role); err != nil {
		global.Logger.Error("Lỗi khi cập nhật role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40004)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"role": role})
}

// DeleteRole @Summary Xóa Role
// @Description Xóa một Role khỏi hệ thống theo ID.
// @Tags Roles
// @Produce json
// @Param id path int true "ID Role cần xóa"
// @Success 200 {object} map[string]interface{} "Xóa role thành công"
// @Failure 40001 {object} map[string]interface{} "ID không hợp lệ"
// @Failure 40005 {object} map[string]interface{} "Lỗi khi xóa role"
// @Router /api/v1/roles/{id} [delete]
func (c *rolesController) DeleteRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	if err := c.roleService.DeleteRole(id); err != nil {
		global.Logger.Error("Lỗi khi xóa role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40005)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"message": "Đã xóa Role", "id": id})
}

// GetPermissionsFromRole @Summary Lấy danh sách permission theo Role
// @Description Lấy tất cả quyền được gán cho một Role cụ thể.
// @Tags Roles
// @Produce json
// @Param id path int true "ID Role"
// @Success 200 {object} map[string]interface{} "Lấy danh sách permissions thành công"
// @Failure 40001 {object} map[string]interface{} "ID không hợp lệ"
// @Failure 40006 {object} map[string]interface{} "Lỗi khi lấy permissions theo role"
// @Router /api/v1/roles/{id}/permissions [get]
func (c *rolesController) GetPermissionsFromRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	permissions, err := c.roleService.GetPermissionsFromRole(id)
	if err != nil {
		global.Logger.Error("Lỗi khi lấy permissions theo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40006)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"permissions": permissions})
}

// AssignPermissionsToRole @Summary Gán quyền cho Role
// @Description Gán một danh sách các quyền cho Role theo ID.
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "ID Role"
// @Param permissionIds body object true "Danh sách ID của quyền được gán"
// @Success 200 {object} map[string]interface{} "Gán permission thành công"
// @Failure 40001 {object} map[string]interface{} "ID không hợp lệ"
// @Failure 40003 {object} map[string]interface{} "Lỗi bind JSON"
// @Failure 40007 {object} map[string]interface{} "Lỗi khi phân quyền role"
// @Router /api/v1/roles/{id}/permissions [post]
func (c *rolesController) AssignPermissionsToRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	var input struct {
		PermissionIDs []uint `json:"permissionIds"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		global.Logger.Error("Lỗi khi bind JSON phân quyền role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	if err := c.roleService.AssignPermissionsToRole(id, input.PermissionIDs); err != nil {
		global.Logger.Error("Lỗi khi phân quyền role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40007)
		return
	}
	response.SuccessResponse(ctx, 20001, gin.H{"message": "Đã gán Permissions cho ", "roleId": id})
}
