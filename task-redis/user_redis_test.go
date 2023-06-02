package user_queue

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// running test queue user to redis
func TestUserRedis(t *testing.T) {
	//initialization redis address and db
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   2,
	})

	ctx := context.Background()

	//calling function NewUserRedisQueue
	redisQueue := NewUserRedisQueue(client)

	//push data user to redis
	redisQueue.PushQueue(ctx, &User{
		ID:       "usr-001",
		Username: "yasir",
		Password: "12345678",
	})

	//push data user to redis
	redisQueue.PushQueue(ctx, &User{
		ID:       "usr-002",
		Username: "nur",
		Password: "12345678",
	})

	//pop data user from redis
	user, err := redisQueue.PopQueue(ctx)
	require.NoError(t, err)

	assert.Equal(t, user.ID, "usr-001")
	assert.Equal(t, user.Username, "yasir")
	assert.Equal(t, user.Password, "12345678")
}
