package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"vietha/src/entity"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "root:12341234@tcp(127.0.0.1:3306)/hagolang?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối database:", err)
	}
	fmt.Println("Kết nối MySQL thành công!")
	DB.AutoMigrate(&entity.User{}, &entity.UserToken{}, &entity.Role{})
	return DB
}
