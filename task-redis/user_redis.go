package user_queue

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

var _ UserQueue = &RedisQueue{}

// create function for NewUserRedisQueue
func NewUserRedisQueue(client *redis.Client) UserQueue {
	return &RedisQueue{client: client}
}

// function count queue
func (r *RedisQueue) Count(ctx context.Context) int64 {
	return r.client.LLen(ctx, "queue").Val()
}

// function push queue to redis
func (r *RedisQueue) PushQueue(ctx context.Context, user *User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.client.LPush(ctx, "queue", bytes).Err()
}

// function pop queue from redis
func (r *RedisQueue) PopQueue(ctx context.Context) (*User, error) {
	result, err := r.client.RPop(ctx, "queue").Result()
	if err != nil {
		return nil, err
	}

	var user User

	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
