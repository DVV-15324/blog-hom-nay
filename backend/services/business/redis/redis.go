package redis

import (
	entityUser "bloghomnay-project/services/entity/user"
	"context"
	"encoding/json"

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

	return bz.redisC.Set(ctx, key, jsonData, 0).Err()
}

// GetProfileEntity trả về struct Users từ Redis
func (bz *BusinessRedis) GetProfileEntity(ctx context.Context, userID string) (*entityUser.Users, error) {
	dataStr, err := bz.redisC.Get(ctx, "profile:"+userID).Result()
	if err != nil {
		return nil, err
	}

	var user entityUser.Users
	if err := json.Unmarshal([]byte(dataStr), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
