package redis

import (
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type BusinessRedis struct {
	redisC *redis.Client
}

func NewBusinessRedis(redisC *redis.Client) *BusinessRedis {
	return &BusinessRedis{
		redisC: redisC,
	}
}

func (bz *BusinessRedis) SaveProfile(ctx context.Context, data *entityUser.Users) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key := "profile:" + data.FakeId

	log.Println("Đang lưu profile vào Redis với key:", key)
	log.Println("Dữ liệu JSON:", string(jsonData))

	return bz.redisC.Set(ctx, key, jsonData, 0).Err()
}

// GetProfile lấy dữ liệu JSON từ Redis theo key "profile:<userID>"
func (bz *BusinessRedis) GetProfile(ctx context.Context, userID string) (string, error) {
	return bz.redisC.Get(ctx, "profile:"+userID).Result()
}
