package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func New(conf Config, timeout time.Duration) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Pass,

		DB: conf.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := client.WithContext(ctx).Ping().Err(); err != nil {
		return nil, fmt.Errorf("redis: ping: %w", err)
	}

	return client, nil
}
