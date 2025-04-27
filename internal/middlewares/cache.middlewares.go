package middlewares

import (
	"awesomeProject5/global"
	"awesomeProject5/response"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // Ghi vào bộ đệm (để cache)
	return w.ResponseWriter.Write(b) // Gửi ra client như bình thường
}
func CacheResponse(duration time.Duration, keyPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := keyPrefix + ":" + c.Request.RequestURI
		// Check Redis
		cachedData, err := global.RedisClient.Get(c, cacheKey).Bytes()
		if err == nil {
			// Cache hit
			c.Data(http.StatusOK, "application/json", cachedData)
			c.Abort()
			return
		}
		// Capture response
		capturingWriter := &bodyWriter{
			ResponseWriter: c.Writer, // giữ bản gốc để vẫn gửi ra client
			body:           bytes.NewBufferString(""),
		}
		c.Writer = capturingWriter
		c.Next()
		// Cache only 200 responses
		if c.Writer.Status() == http.StatusOK {
			_ = global.RedisClient.Set(c, cacheKey, capturingWriter.body.Bytes(), duration).Err()
		}
	}
}
func CacheResponseWithKeyFunc(duration time.Duration, keyFunc func(*gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := keyFunc(c)
		cachedData, err := global.RedisClient.Get(c, cacheKey).Bytes()
		if err == nil {
			c.Data(http.StatusOK, "application/json", cachedData)
			c.Abort()
			return
		}
		capturingWriter := &bodyWriter{
			ResponseWriter: c.Writer, // giữ bản gốc để vẫn gửi ra client
			body:           bytes.NewBufferString(""),
		}
		c.Writer = capturingWriter
		c.Next()
		var resp response.ResponseData
		if err := json.Unmarshal(capturingWriter.body.Bytes(), &resp); err == nil && resp.Code == 20001 {
			_ = global.RedisClient.Set(c, cacheKey, capturingWriter.body.Bytes(), duration).Err()
		}
	}
}
func InvalidateCache(patternFunc func(c *gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ghi đè Writer để capture nội dung phản hồi
		capturingWriter := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = capturingWriter

		c.Next()

		var resp response.ResponseData
		if err := json.Unmarshal(capturingWriter.body.Bytes(), &resp); err == nil && resp.Code == 20001 {
			ctx := context.Background()
			pattern := patternFunc(c)
			iter := global.RedisClient.Scan(ctx, 0, pattern, 0).Iterator()
			for iter.Next(ctx) {
				key := iter.Val()
				if err := global.RedisClient.Del(ctx, key).Err(); err != nil {
					global.Logger.Error("Lỗi khi xóa cache key %s: %v")
				}
			}
			if err := iter.Err(); err != nil {
				global.Logger.Error("Lỗi khi quét Redis keys: %v")
			}
		}
	}
}
