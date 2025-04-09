package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vietha/src/entity"
)

type PermissionController interface {
	GetAllPermissions(ctx *gin.Context)
	GetPermissionById(ctx *gin.Context)
	CreatePermission(ctx *gin.Context)
	UpdatePermission(ctx *gin.Context)
	DeletePermission(ctx *gin.Context)
}
type permissionController struct {
	db *gorm.DB
}

func NewPermissionController(db *gorm.DB) PermissionController {
	return &permissionController{
		db: db,
	}
}
func (c *permissionController) GetAllPermissions(ctx *gin.Context) {
	var permissions []entity.Permissions
	c.db.Find(&permissions)
	ctx.JSON(200, permissions)
}

func (c *permissionController) GetPermissionById(ctx *gin.Context) {
	id := ctx.Param("id")
	var permission entity.Permissions
	if err := c.db.First(&permission, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Permission not found"})
		return
	}
	ctx.JSON(200, permission)
}

func (c *permissionController) CreatePermission(ctx *gin.Context) {
	var permission entity.Permissions
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.db.Create(&permission)
	ctx.JSON(201, permission)
}

func (c *permissionController) UpdatePermission(ctx *gin.Context) {
	id := ctx.Param("id")
	var permission entity.Permissions
	if err := c.db.First(&permission, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Permission not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.db.Save(&permission)
	ctx.JSON(200, permission)
}

func (c *permissionController) DeletePermission(ctx *gin.Context) {
	id := ctx.Param("id")
	c.db.Delete(&entity.Permissions{}, id)
	ctx.JSON(200, gin.H{"message": "Permission deleted"})
}
