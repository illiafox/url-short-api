package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(timeout time.Duration, cfg Config) (*pgxpool.Pool, error) {

	ctx := context.Background()
	if timeout > 0 {
		wt, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		ctx = wt
	}

	return pgxpool.Connect(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%v:%v/%v?sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.IP,
			cfg.Port,
			cfg.DBName,
		),
	)
}
