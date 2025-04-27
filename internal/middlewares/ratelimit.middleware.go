package middlewares

import (
	"awesomeProject5/global"
	"context"
	"fmt"
	"time"
)

// RateLimit kiểm tra xem hệ thống có vượt quá giới hạn request trong khoảng thời gian cụ thể không
func RateLimit(limit int, window time.Duration) (bool, int64) {
	ctx := context.Background()
	// Key dùng chung cho toàn bộ hệ thống, theo từng "window"
	key := fmt.Sprintf("rate_limit:global:%d", time.Now().Unix()/int64(window.Seconds()))
	// Tăng số lượng request trong Redis
	count, err := global.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Lỗi Redis:", err)
		return false, 0
	}
	// Nếu là request đầu tiên trong window, đặt expire
	if count == 1 {
		global.RedisClient.Expire(ctx, key, window)
	}
	return count <= int64(limit), count
}
