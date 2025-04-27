package initialize

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/mock"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitRedis() {
	env := global.Config.App.Env
	if env == "dev" {
		global.RedisClient = redis.NewClient(&redis.Options{
			Addr:         "localhost:6379", // địa chỉ Redis server
			Password:     "",
			DB:           0,
			PoolSize:     100,             // Số lượng kết nối tối đa trong pool
			MinIdleConns: 10,              // Số kết nối tối thiểu duy trì
			DialTimeout:  5 * time.Second, // Thời gian timeout khi tạo kết nối
			ReadTimeout:  3 * time.Second, // Thời gian timeout khi đọc
			WriteTimeout: 3 * time.Second, // Thời gian timeout khi ghi
			PoolTimeout:  4 * time.Second, // Thời gian timeout khi lấy kết nối từ pool
		})
		global.Logger.Info("Dùng Redis Log")
	} else {
		global.RedisClient = &mock.NullRedisClient{}
		global.Logger.Info("Dùng Redis Mock")
	}
	_, err := global.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Không thể kết nối Redis: " + err.Error())
	} else {
		global.Logger.Info("Đã kết nối Redis")
	}

}
