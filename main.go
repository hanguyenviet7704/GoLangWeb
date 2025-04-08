package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
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
	db                 *gorm.DB                      = config.ConnectDatabase()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	userService        service.UserService           = service.NewUserService(userRepository, db)
	jwtService         service.JWTService            = service.JWTAuthService(db)
	loginController    controller.LoginController    = controller.LoginHandler(db, userService, jwtService)
	logoutController   controller.LogoutController   = controller.NewLogout(db)
	registerController controller.RegisterController = controller.NewRegisterController(userService)
	authHandlers                                     = handlers.NewAuth(db)
	userController                                   = controller.NewUserController(userService)
)

// @title Nông sản API
// @version 1.0
// @description API cho hệ thống bán hàng nông sản trực tuyến
// @host localhost:2004
// @BasePath /api
func main() {
	// Khởi tạo các controller và dịch vụ
	server := gin.New()
	// Users API
	server.GET("api/users", userController.GetAllUsers)                  // Lấy tất cả người dùng
	server.GET("api/user", userController.GetUserById)                   // Lấy thông tin người dùng theo ID
	server.POST("api/users", userController.CreateOrUpdateUser)          // Tạo hoặc cập nhật người dùng
	server.DELETE("/api/users", userController.DeleteUser)               // Xóa người dùng
	server.GET("/api/users/:id/roles", userController.GetRolesFromUser)  // Lấy quyền của người dùng
	server.PUT("/api/users/:id/roles", userController.AssignRolesToUser) // Gán quyền cho người dùng
	// Authentication API
	server.POST("api/auth/login", loginController.Login)          // Đăng nhập
	server.POST("api/auth/refresh", loginController.RefreshToken) // Làm mới token
	server.POST("api/auth/register", func(ctx *gin.Context) {
		registerController.Register(ctx) // Đăng ký người dùng
	})
	server.POST("api/auth/logout", logoutController.Logout)                     // Đăng xuất
	server.POST("api/auth/logoutalldevices", logoutController.LogoutAllDevices) // Đăng xuất khỏi tất cả thiết bị
	// Swagger UI
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Nhóm route bảo vệ bởi JWT
	authGroup := server.Group("/api")
	authGroup.Use(middleware.AuthorizeJWT(jwtService)) // Middleware kiểm tra JWT
	{
		authGroup.GET("/protected", func(c *gin.Context) {
			// Lấy claims từ context
			claims, exists := c.Get("claims")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				return
			}
			// Kiểm tra kiểu dữ liệu của claims
			authClaims, ok := claims.(*service.AuthCustomClaims)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				return
			}
			// Trả về dữ liệu claims
			c.JSON(http.StatusOK, gin.H{
				"message":     "This is validate token",
				"email":       authClaims.Email,
				"name":        authClaims.Name,
				"userid":      authClaims.UserID,
				"issuer":      authClaims.Issuer,
				"expires":     authClaims.ExpiresAt,
				"roles":       authClaims.Roles,
				"permissions": authClaims.Permissions,
			})
		})
	}
	//
	// Khởi động server
	port := os.Getenv("PORT")
	if port == "" {
		port = "2004"
	}
	err := server.Run(":" + port)
	if err != nil {
		panic("Không thể khởi động server: " + err.Error())
	}
}

//http://localhost:2004/api/protected
