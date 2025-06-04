# GOLANDWEB API Documentation

## Giới thiệu
GOLANDWEB API là một RESTful API được xây dựng bằng Go, cung cấp các chức năng xác thực và phân quyền người dùng.

## Cài đặt và Chạy

### Yêu cầu hệ thống
- Go 1.16 trở lên
- MySQL/PostgreSQL
- Make (tùy chọn)

### Cài đặt
1. Clone repository:
```bash
git clone <repository-url>
cd GOLANDWEB
```

2. Cài đặt dependencies:
```bash
go mod download
```

3. Cấu hình môi trường:
- Tạo file `.env` từ file `.env.example`
- Cập nhật các biến môi trường cần thiết

4. Chạy migrations:
```bash
go run migrations/migrate.go
```

5. Chạy ứng dụng:
```bash
go run main.go
```

## API Endpoints

### Authentication

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password123"
}
```

Response (200 OK):
```json
{
    "token": "jwt-token-here",
    "user": {
        "id": 1,
        "email": "user@example.com",
        "role_id": 1
    }
}
```

#### Logout
```http
POST /api/v1/auth/logout
Authorization: Bearer <token>
```

#### Logout All Devices
```http
POST /api/v1/auth/logout-all
Authorization: Bearer <token>
Content-Type: application/json

{
    "user_id": 1
}
```

#### Refresh Token
```http
POST /api/v1/auth/refresh
Authorization: Bearer <refresh-token>
```

### Users

#### Get All Users
```http
GET /api/v1/users
Authorization: Bearer <token>
```

Response (200 OK):
```json
{
    "users": [
        {
            "id": 1,
            "email": "user@example.com",
            "role_id": 1,
            "created_at": "2024-03-14T10:00:00Z"
        }
    ]
}
```

#### Get User by ID
```http
GET /api/v1/users/{id}
Authorization: Bearer <token>
```

#### Delete User
```http
DELETE /api/v1/users/{id}
Authorization: Bearer <token>
```

### Roles

#### Get All Roles
```http
GET /api/v1/roles
Authorization: Bearer <token>
```

Response (200 OK):
```json
{
    "roles": [
        {
            "id": 1,
            "name": "Admin",
            "description": "Administrator role",
            "permissions": [
                {
                    "id": 1,
                    "name": "read",
                    "action": "GET",
                    "resource": "users"
                }
            ]
        }
    ]
}
```

#### Create Role
```http
POST /api/v1/roles
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "New Role",
    "description": "Role description"
}
```

#### Update Role
```http
PUT /api/v1/roles/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "Updated Role",
    "description": "Updated description"
}
```

#### Delete Role
```http
DELETE /api/v1/roles/{id}
Authorization: Bearer <token>
```

### Role Permissions

#### Get Role Permissions
```http
GET /api/v1/roles/{id}/permissions
Authorization: Bearer <token>
```

#### Assign Permissions to Role
```http
POST /api/v1/roles/{id}/permissions
Authorization: Bearer <token>
Content-Type: application/json

{
    "permissionIds": [1, 2, 3]
}
```

## Mã lỗi

| Mã lỗi | Mô tả |
|--------|--------|
| 40001 | Lỗi hệ thống |
| 40002 | Lỗi khi truy vấn |
| 40003 | Lỗi bind JSON |
| 40004 | Lỗi tạo/cập nhật |
| 40005 | Lỗi khi xóa |
| 40006 | Lỗi khi lấy permissions |
| 40007 | Lỗi khi phân quyền role |
| 50001 | Không tìm thấy người dùng |
| 50003 | Sai mật khẩu |

## Bảo mật

- Tất cả các API endpoints (trừ login) yêu cầu JWT token trong header
- Format token: `Authorization: Bearer <token>`
- Token có thời hạn sử dụng và có thể được refresh
- Mỗi role có các quyền riêng biệt

## Testing

### Chạy tests
```bash
# Chạy tất cả tests
go test ./tests

# Chạy test với verbose output
go test -v ./tests

# Chạy test cụ thể
go test -v ./tests -run TestLoginAPI
```

### Test Coverage
```bash
# Tạo coverage report
go test -coverprofile=coverage.out ./tests
go tool cover -html=coverage.out
```

## Swagger Documentation

Truy cập Swagger UI tại: http://localhost:8080/swagger/index.html

## Liên hệ

- Tác giả: Việt Hà
- Email: hanguyenviet772004@gmail.com
