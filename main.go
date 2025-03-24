package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
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
	userService        service.UserService           = service.NewUserService(userRepository)
	jwtService         service.JWTService            = service.JWTAuthService()
	loginController    controller.LoginController    = controller.LoginHandler(db, userService, jwtService)
	logoutController   controller.LogoutController   = controller.NewLogout(db)
	registerController controller.RegisterController = controller.NewRegisterController(userService)
	authHandlers                                     = handlers.NewAuth(db)
)

func main() {
	server := gin.New()
	server.POST("/login", loginController.Login)
	server.POST("/refresh", loginController.RefreshToken)
	server.POST("/register", func(ctx *gin.Context) {
		registerController.Register(ctx) // Gọi đúng RegisterController
	})
	server.POST("/logout", logoutController.Logout)
	server.POST("/logoutalldevices", logoutController.LogoutAllDevices)
	server.POST("/logoutdevice", logoutController.LogoutDevice)
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
				"message": "Bạn đã truy cập vào!",
				"email":   authClaims.Email,
				"user":    authClaims.User,
			})
		})
	}
	server.POST("/forgot-password", authHandlers.ForgotPassword)
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
