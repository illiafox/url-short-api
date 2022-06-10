package links

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"ozon-url-shortener/app/internal/domain/links"
)

type PgStorage struct {
	pool *pgxpool.Pool
}

func NewPgStorage(pool *pgxpool.Pool) links.Storage {
	return PgStorage{
		pool: pool,
	}
}

func (s PgStorage) GetURL(ctx context.Context, key string) (url string, err error) {

	err = s.pool.
		QueryRow(ctx, "SELECT url FROM links WHERE key = $1", key).
		Scan(&url)

	return
}

func (s PgStorage) StoreURL(ctx context.Context, key []byte, url string) error {
	_, err := s.pool.Exec(ctx, "INSERT INTO links VALUES ($1,$2)", key, url)

	return err
}
