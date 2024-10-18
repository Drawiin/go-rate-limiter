package infra

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	ctx := context.Background()
	err := client.Set(ctx, "testkey", "testvalue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "testkey").Result()
	if err != nil || val != "testvalue" {
		panic("failed to verify Redis client")
	}

	return &RedisStore{client: client}
}

func (r *RedisStore) Set(key string, timeWindow, requestCount int64) {
	ctx := context.Background()
	r.client.HSet(ctx, key, "timeWindow", timeWindow)
	r.client.HSet(ctx, key, "requestCount", requestCount)
}

func (r *RedisStore) Get(key string) (int64, int64) {
	ctx := context.Background()
	values, err := r.client.HGetAll(ctx, key).Result()
	if err != nil || len(values) == 0 {
		return 0, 0
	}

	timeWindow, _ := strconv.ParseInt(values["timeWindow"], 10, 64)
	requestCount, _ := strconv.ParseInt(values["requestCount"], 10, 64)
	return timeWindow, requestCount
}