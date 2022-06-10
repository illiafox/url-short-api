package links

import "context"

type Storage interface {
	GetURL(ctx context.Context, key string) (url string, err error)
	StoreURL(ctx context.Context, key []byte, url string) error
}
