package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vietha/src/entity"
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
	db *gorm.DB
}

func NewRolesController(db *gorm.DB) RolesController {
	return &rolesController{
		db: db,
	}
}

func (c *rolesController) GetAllRoles(ctx *gin.Context) {
	var roles []entity.Roles
	c.db.Preload("Permissions").Find(&roles)
	ctx.JSON(200, roles)
}

func (c *rolesController) GetRoleById(ctx *gin.Context) {
	id := ctx.Param("id")
	var role entity.Roles
	if err := c.db.Preload("Permissions").First(&role, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(200, role)
}

func (c *rolesController) CreateRole(ctx *gin.Context) {
	
	var role entity.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.db.Create(&role)
	ctx.JSON(201, role)
}

func (c *rolesController) UpdateRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var role entity.Roles
	if err := c.db.First(&role, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.db.Save(&role)
	ctx.JSON(200, role)
}

func (c *rolesController) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")
	c.db.Delete(&entity.Roles{}, id)
	ctx.JSON(200, gin.H{"message": "Role deleted"})
}

func (c *rolesController) GetPermissionsFromRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var role entity.Roles
	if err := c.db.Preload("Permissions").First(&role, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(200, role.Permissions)
}

func (c *rolesController) AssignPermissionsToRole(ctx *gin.Context) {
	id := ctx.Param("id")
	var permissionIDs []uint
	if err := ctx.ShouldBindJSON(&permissionIDs); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var role entity.Roles
	if err := c.db.First(&role, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Role not found"})
		return
	}
	var permissions []entity.Permissions
	c.db.Where("id IN (?)", permissionIDs).Find(&permissions)
	c.db.Model(&role).Association("Permissions").Replace(permissions)
	ctx.JSON(200, gin.H{"message": "Permissions assigned to role"})
}
