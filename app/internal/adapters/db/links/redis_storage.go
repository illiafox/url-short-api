package links

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"ozon-url-shortener/app/internal/domain/links"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(client *redis.Client) links.Storage {
	return RedisStorage{
		client: client,
	}
}

func (s RedisStorage) GetURL(ctx context.Context, key string) (string, error) {
	url, err := s.client.WithContext(ctx).Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("redis: get: %w", err)
	}

	return url, nil
}

const expire = time.Hour * 24 * 30

func (s RedisStorage) StoreURL(ctx context.Context, key []byte, url string) error {
	err := s.client.WithContext(ctx).Set(string(key), url, expire).Err()
	if err != nil {
		return fmt.Errorf("redis: set: %w", err)
	}

	return nil
}
