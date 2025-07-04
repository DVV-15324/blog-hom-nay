package composer

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// InitRedis khởi tạo Redis client 1 lần duy nhất
func InitRedis(addr, password string, db int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}
	log.Println("Redis connected thành công")
	return redisClient
}
