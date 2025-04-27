package initialize

import (
	"awesomeProject5/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func InitRouter(app *AppContainer) *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())
	// ========================= Swagger router =========================
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ========================= Auth APIs =========================
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", app.LoginController.Login)
		authRoutes.POST("/refresh", app.LoginController.RefreshToken)
		authRoutes.POST("/register", app.RegisterController.Register)
		// Các route sau cần xác thực bằng JWT
		authRoutes.Use(app.MiddlewareAuth.AuthorizeJWT())
		authRoutes.POST("/logout", app.LogoutController.Logout)
		authRoutes.POST("/logoutalldevices", app.LogoutController.LogoutAllDevices)
	}
	// ========================= Admin APIs =========================
	adminRoutes := server.Group("/api/admin")
	adminRoutes.Use(
		app.MiddlewareAuth.AuthorizeJWT(),
		app.MiddlewareAuth.RoleMiddleware("ROLE_ADMIN"),
	)
	{
		// Users
		// Lấy tất cả user
		adminRoutes.GET("/users",
			middlewares.CacheResponse(5*time.Minute, "admin:users"),
			app.UserController.GetAllUsers)
		// Lấy user theo id
		adminRoutes.GET("/users/:id",
			middlewares.CacheResponseWithKeyFunc(5*time.Minute, func(c *gin.Context) string {
				return "admin:users:" + c.Param("id")
			}),
			app.UserController.GetUserById)
		// Thêm hoạc cập nhật user
		adminRoutes.POST("/users",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:users:" + c.Param("id")
			}),
			app.UserController.CreateOrUpdateUser)
		// Xóa user
		adminRoutes.DELETE("/users",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:users:" + c.Param("id")
			}),
			app.UserController.DeleteUser)
		// Lấy user theo
		adminRoutes.GET("/users/:id/roles",
			middlewares.CacheResponseWithKeyFunc(5*time.Minute, func(c *gin.Context) string {
				return "admin:users:" + c.Param("id") + ":roles"
			}),
			app.UserController.GetRolesFromUser)
		// Thêm quyền cho user
		adminRoutes.PUT("/users/:id/roles",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:users:" + c.Param("id")
			}),
			app.UserController.AssignRolesToUser)
		// Roles
		adminRoutes.GET("/roles",
			middlewares.CacheResponse(5*time.Minute, "admin:roles"),
			app.RoleController.GetAllRoles)
		adminRoutes.GET("/roles/:id",
			middlewares.CacheResponseWithKeyFunc(5*time.Minute, func(c *gin.Context) string {
				return "admin:roles:" + c.Param("id")
			}),
			app.RoleController.GetRoleById)

		adminRoutes.POST("/roles",
			app.RoleController.CreateRole)
		adminRoutes.PUT("/roles/:id",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:roles:" + c.Param("id")
			}),
			app.RoleController.UpdateRole)
		adminRoutes.DELETE("/roles/:id",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:roles:" + c.Param("id")
			}),
			app.RoleController.DeleteRole)
		adminRoutes.GET("/roles/:id/permissions",
			middlewares.CacheResponseWithKeyFunc(5*time.Minute, func(c *gin.Context) string {
				return "admin:roles:" + c.Param("id") + ":permissons"
			}),
			app.RoleController.GetPermissionsFromRole)
		adminRoutes.PUT("/roles/:id/permissions", app.RoleController.AssignPermissionsToRole)
		// Permissions
		adminRoutes.GET("/permissions",
			middlewares.CacheResponse(5*time.Minute, "admin:permissons"),
			app.PermissionController.GetAllPermissions)
		adminRoutes.GET("/permissions/:id",
			middlewares.CacheResponseWithKeyFunc(5*time.Minute, func(c *gin.Context) string {
				return "admin:permissons:" + c.Param("id")
			}),
			app.PermissionController.GetPermissionById)
		adminRoutes.POST("/permissions",

			app.PermissionController.CreatePermission)
		adminRoutes.PUT("/permissions/:id",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:permissions:" + c.Param("id")
			}),
			app.PermissionController.UpdatePermission)
		adminRoutes.DELETE("/permissions/:id",
			middlewares.InvalidateCache(func(c *gin.Context) string {
				return "admin:permissions:" + c.Param("id")
			}),
			app.PermissionController.DeletePermission)
	}
	// ========================= Protected APIs =========================
	protected := server.Group("/api")
	protected.Use(app.MiddlewareAuth.AuthorizeJWT())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Thông tin người dùng đã đăng nhập"})
		})
		protected.GET("/products",
			app.MiddlewareAuth.PermissionMiddleware("view"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Bạn có quyền xem danh sách sản phẩm"})
			},
		)
		protected.GET("/dashboard",
			app.MiddlewareAuth.RoleMiddleware("ROLE_ADMIN", "ROLE_USER"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Chào bạn, đây là dashboard dành cho admin hoặc manager"})
			},
		)
	}

	return server
}
