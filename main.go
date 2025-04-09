package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"os"
	_ "vietha/docs"
	"vietha/src/config"
	"vietha/src/controller"
	"vietha/src/handlers"
	"vietha/src/middleware"
	"vietha/src/repository"
	"vietha/src/service"
)

var (
	db                   *gorm.DB                      = config.ConnectDatabase()
	userRepository       repository.UserRepository     = repository.NewUserRepository(db)
	userService          service.UserService           = service.NewUserService(userRepository, db)
	jwtService           service.JWTService            = service.JWTAuthService(db)
	loginController      controller.LoginController    = controller.LoginHandler(db, userService, jwtService)
	logoutController     controller.LogoutController   = controller.NewLogout(db)
	registerController   controller.RegisterController = controller.NewRegisterController(userService)
	middlewareAuth                                     = middleware.MiddlewareHandler(jwtService)
	auth                                               = handlers.NewAuth(db, middlewareAuth, jwtService)
	userController                                     = controller.NewUserController(userService)
	roleController                                     = controller.NewRolesController(db)
	permissionController                               = controller.NewPermissionController(db)
)

// @title Nông sản API
// @version 1.0
// @description API cho hệ thống bán hàng nông sản trực tuyến
// @host localhost:2004
// @BasePath /api
func main() {
	// Tạo router mới với Logger và Recovery
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())
	// ========================= Swagger =========================
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ========================= Auth APIs =========================
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", loginController.Login)
		authRoutes.POST("/refresh", loginController.RefreshToken)
		authRoutes.POST("/register", func(ctx *gin.Context) {
			registerController.Register(ctx)
		})
		// Các route sau cần xác thực bằng JWT
		authRoutes.Use(middlewareAuth.AuthorizeJWT())
		authRoutes.POST("/logout", logoutController.Logout)
		authRoutes.POST("/logoutalldevices", logoutController.LogoutAllDevices)
	}
	// ========================= Admin APIs =========================
	adminRoutes := server.Group("/api/admin")
	adminRoutes.Use(middlewareAuth.AuthorizeJWT(), middlewareAuth.RoleMiddleware("ROLE_ADMIN"))
	{
		// Users
		adminRoutes.GET("/users", userController.GetAllUsers)
		adminRoutes.GET("/user", userController.GetUserById)
		adminRoutes.POST("/users", userController.CreateOrUpdateUser)
		adminRoutes.DELETE("/users", userController.DeleteUser)
		adminRoutes.GET("/users/:id/roles", userController.GetRolesFromUser)
		adminRoutes.PUT("/users/:id/roles", userController.AssignRolesToUser)
		// Roles
		adminRoutes.GET("/roles", roleController.GetAllRoles)
		adminRoutes.GET("/roles/:id", roleController.GetRoleById)
		adminRoutes.POST("/roles", roleController.CreateRole)
		adminRoutes.PUT("/roles/:id", roleController.UpdateRole)
		adminRoutes.DELETE("/roles/:id", roleController.DeleteRole)
		adminRoutes.GET("/roles/:id/permissions", roleController.GetPermissionsFromRole)
		adminRoutes.PUT("/roles/:id/permissions", roleController.AssignPermissionsToRole)
		// Permissions
		adminRoutes.GET("/permissions", permissionController.GetAllPermissions)
		adminRoutes.GET("/permissions/:id", permissionController.GetPermissionById)
		adminRoutes.POST("/permissions", permissionController.CreatePermission)
		adminRoutes.PUT("/permissions/:id", permissionController.UpdatePermission)
		adminRoutes.DELETE("/permissions/:id", permissionController.DeletePermission)
	}
	// ========================= Protected APIs =========================
	protected := server.Group("/api")
	protected.Use(middlewareAuth.AuthorizeJWT())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Thông tin người dùng đã đăng nhập"})
		})
		protected.GET("/products", middlewareAuth.PermissionMiddleware("view_products"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Bạn có quyền xem danh sách sản phẩm"})
		})
		protected.GET("/dashboard",
			middlewareAuth.RoleMiddleware("ROLE_ADMIN", "ROLE_MANAGER"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Chào bạn, đây là dashboard dành cho admin hoặc manager"})
			},
		)
	}
	// ========================= Start Server =========================
	port := os.Getenv("PORT")
	if port == "" {
		port = "2004"
	}
	if err := server.Run(":" + port); err != nil {
		panic("Không thể khởi động server: " + err.Error())
	}
}
