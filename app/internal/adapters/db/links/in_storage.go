package links

import (
	"context"

	"ozon-url-shortener/app/internal/domain/links"
	"ozon-url-shortener/app/pkg/cache"
)

type MemStorage struct {
	cache cache.Cache
}

func NewMemStorage() links.Storage {
	return &MemStorage{
		cache: cache.New(),
	}
}

func (m *MemStorage) GetURL(ctx context.Context, key string) (url string, err error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return m.cache.Get(key), nil
	}
}

func (m *MemStorage) StoreURL(ctx context.Context, key []byte, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		m.cache.Set(string(key), url)

		return nil
	}
}
