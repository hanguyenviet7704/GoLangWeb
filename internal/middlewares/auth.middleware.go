package middlewares

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sort"
	"strings"
	"time"
)

type JWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
	RoleMiddleware(allowedRoles ...string) gin.HandlerFunc
	PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc
}
type jwtMiddleware struct {
	jwtService  service.JWTService
	userService service.UserService
}

func NewJWTMiddleware(jwtService service.JWTService, userService service.UserService) JWTMiddleware {
	return &jwtMiddleware{
		jwtService:  jwtService,
		userService: userService,
	}
}

// Thằng này nếu dùng cache vào thì chậm hơn chắc là do có các phần Marshal và Unmarshal
func (m *jwtMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			global.Logger.Warn("Authorization header rỗng")
			response.ErrorResponse(ctx, 20002)
			ctx.Abort()
			return
		}
		const BEARER_SCHEMA = "Bearer "
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			global.Logger.Warn("Lỗi định dạng Authorization")
			response.ErrorResponse(ctx, 20002)
			ctx.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		// 1. Kiểm tra token trong Redis cache
		cacheKey := fmt.Sprintf("jwt:claims:%s", tokenString)
		cachedClaims, err := global.RedisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			// Nếu tồn tại, parse lại claims và bỏ qua bước validate
			var claims *service.AuthCustomClaims
			if err := json.Unmarshal([]byte(cachedClaims), &claims); err == nil {
				ctx.Set("claims", claims)
				ctx.Next()
				return
			}
			// Nếu lỗi parse thì tiếp tục validate token như bình thường
		}
		// 2. Nếu không có trong cache, validate token
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			global.Logger.Error("Lỗi xác thực token", zap.Error(err))
			response.ErrorResponse(ctx, 30001)
			ctx.Abort()
			return
		}
		// 3. Cache lại claims (convert sang JSON để lưu)
		if jsonBytes, err := json.Marshal(claims); err == nil {
			_ = global.RedisClient.Set(ctx, cacheKey, jsonBytes, time.Minute*15).Err() // cache 15 phút
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

// Thằng này thì lưu user cộng với vai trò  nếu có thì oke luôn . Ở đây nếu không dùng redis chạy khoảng 13s vì mỗi lần sẽ check quyền trong database . còn nếu dùng thì sẽ chạy khoảng 0.45s
func (m *jwtMiddleware) RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			global.Logger.Warn("Không tìm thấy claims trong context")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			global.Logger.Error("Kiểu dữ liệu claims không hợp lệ")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		// Sort allowedRoles để đảm bảo tính nhất quán của cache key
		sort.Strings(allowedRoles)
		roleKey := fmt.Sprintf("user:%d:roles:%s", authClaims.UserID, strings.Join(allowedRoles, "_"))
		// Kiểm tra cache
		hasRole, err := global.RedisClient.Get(ctx, roleKey).Bool()
		if err == nil {
			if hasRole {
				ctx.Next()
			} else {
				global.Logger.Warn("Không có vai trò phù hợp để truy cập (cached)", zap.Strings("required_roles", allowedRoles))
				response.ErrorResponse(ctx, 30004)
				ctx.Abort()
			}
			return
		}
		// Cache miss, truy vấn DB
		user, err := m.userService.FindUserWithRolesAndPermissionsByID(authClaims.UserID)
		if err != nil {
			global.Logger.Error("Không thể tìm thấy thông tin người dùng", zap.Error(err))
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		// Kiểm tra vai trò
		for _, userRole := range user.Roles {
			for _, allowed := range allowedRoles {
				if allowed == userRole.Name {
					// Lưu kết quả tích cực vào cache
					_ = global.RedisClient.Set(ctx, roleKey, true, time.Minute*15).Err()
					ctx.Next()
					return
				}
			}
		}
		// Lưu kết quả tiêu cực vào cache (với TTL ngắn hơn)
		_ = global.RedisClient.Set(ctx, roleKey, false, time.Minute*5).Err()
		global.Logger.Warn("Không có vai trò phù hợp để truy cập", zap.Strings("required_roles", allowedRoles))
		response.ErrorResponse(ctx, 30004)
		ctx.Abort()
	}
}

// Thằng này thì lưu user cộng với quyền  nếu có thì oke luôn . Ở đây nếu không dùng redis chạy khoảng 16s vì mỗi lần sẽ check quyền trong database . còn nếu dùng thì sẽ chạy khoảng 0.4s
func (m *jwtMiddleware) PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			global.Logger.Error("Không tìm thấy claims trong context")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			global.Logger.Error("Kiểu dữ liệu claims không hợp lệ")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		// Sort allowedPermissions để đảm bảo tính nhất quán của cache key
		sort.Strings(allowedPermissions)
		permissionKey := fmt.Sprintf("user:%d:permission:%s", authClaims.Id, strings.Join(allowedPermissions, "_"))
		// Kiểm tra cache
		hasPermission, err := global.RedisClient.Get(ctx, permissionKey).Bool()
		if err == nil {
			if hasPermission {
				ctx.Next()
			} else {
				global.Logger.Warn("Không có quyền phù hợp để truy cập (cached)", zap.Strings("required_permissions", allowedPermissions))
				response.ErrorResponse(ctx, 30004)
				ctx.Abort()
			}
			return
		}
		// Cache miss, truy vấn DB
		user, err := m.userService.FindUserWithRolesAndPermissionsByID(authClaims.UserID)
		if err != nil {
			global.Logger.Error("Không thể tìm thấy thông tin người dùng", zap.Error(err))
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		// Tạo một map các quyền từ vai trò
		permMap := make(map[string]bool)
		// Thêm quyền từ vai trò vào map
		for _, userRole := range user.Roles {
			for _, permission := range userRole.Permissions {
				permMap[permission.Action] = true
			}
		}
		// Thêm quyền trực tiếp của user vào map
		for _, userPerm := range user.Permissions {
			permMap[userPerm.Action] = true
		}
		// Kiểm tra nếu có bất kỳ quyền nào trong allowedPermissions
		for _, allowedPermission := range allowedPermissions {
			if permMap[allowedPermission] {
				// Lưu kết quả tích cực vào cache
				_ = global.RedisClient.Set(ctx, permissionKey, true, time.Minute*15).Err()
				ctx.Next()
				return
			}
		}
		// Lưu kết quả tiêu cực vào cache (với TTL ngắn hơn)
		_ = global.RedisClient.Set(ctx, permissionKey, false, time.Minute*5).Err()
		global.Logger.Warn("Không có quyền phù hợp để truy cập", zap.Strings("required_permissions", allowedPermissions))
		response.ErrorResponse(ctx, 30004)
		ctx.Abort()
	}
}
