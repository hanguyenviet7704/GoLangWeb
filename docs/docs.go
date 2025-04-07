// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Hỗ trợ API",
            "email": "support@nongsan.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/users": {
            "get": {
                "description": "Trả về danh sách tất cả người dùng trong hệ thống\nKết quả trả về bao gồm thông tin cơ bản của người dùng và trạng thái hoạt động",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Lấy danh sách người dùng",
                "responses": {
                    "200": {
                        "description": "Danh sách người dùng",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "users": {
                                    "type": "array",
                                    "items": {
                                        "$ref": "#/definitions/entity.User"
                                    }
                                }
                            }
                        },
                        "examples": {
                            "application/json": {
                                "users": [
                                    {
                                        "id": 1,
                                        "name": "Nguyễn Văn A",
                                        "email": "a@nongsan.com",
                                        "is_active": true,
                                        "created_at": "2023-01-01T00:00:00Z",
                                        "updated_at": "2023-01-01T00:00:00Z"
                                    },
                                    {
                                        "id": 2,
                                        "name": "Trần Thị B",
                                        "email": "b@nongsan.com",
                                        "is_active": false,
                                        "created_at": "2023-01-02T00:00:00Z",
                                        "updated_at": "2023-01-02T00:00:00Z"
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "Lỗi hệ thống",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        },
                        "examples": {
                            "application/json": {
                                "error": "Không thể kết nối đến cơ sở dữ liệu"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Permissions": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string",
                    "example": "read"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Xem danh sách người dùng"
                },
                "resource": {
                    "type": "string",
                    "example": "users"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                }
            }
        },
        "entity.Roles": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Quản trị viên"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Permissions"
                    }
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                }
            }
        },
        "entity.Tokens": {
            "type": "object",
            "properties": {
                "access_token_expires_at": {
                    "type": "string",
                    "example": "2023-01-01T01:00:00Z"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "refresh_token": {
                    "type": "string",
                    "example": "abcdef123456"
                },
                "refresh_token_expires_at": {
                    "type": "string",
                    "example": "2023-01-08T00:00:00Z"
                },
                "token": {
                    "type": "string",
                    "example": "xyz123"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "type": "string",
                    "example": "Nguyễn Văn A"
                },
                "password": {
                    "type": "string",
                    "example": "hashedpassword123"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Permissions"
                    }
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Roles"
                    }
                },
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Tokens"
                    }
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2004",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Nông sản API",
	Description:      "API cho hệ thống bán hàng nông sản trực tuyến",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
