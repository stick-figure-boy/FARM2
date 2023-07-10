package redis

import (
	"encoding/json"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Get(ctx *gin.Context, key string) (string, error)
	Set(ctx *gin.Context, key string, value any, expr time.Duration) error
	Delete(ctx *gin.Context, key string) error
	Expire(ctx *gin.Context, key string, expr time.Duration) error
}

type Redis struct {
	client *redis.Client
}

func NewRedisClient() RedisClient {
	return &Redis{
		client: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_DNS"),
			DB:   0,
		}),
	}
}

func (r *Redis) Get(ctx *gin.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx *gin.Context, key string, value any, expr time.Duration) error {
	if reflect.TypeOf(value).Kind() != reflect.String {
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = string(b)
	}
	return r.client.Set(ctx, key, value, expr).Err()
}

func (r *Redis) Delete(ctx *gin.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) Expire(ctx *gin.Context, key string, expr time.Duration) error {
	err := r.client.Expire(ctx, key, expr).Err()
	if err != nil {
		return err
	}
	return nil
}
