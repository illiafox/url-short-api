package app

import (
	"time"

	"go.uber.org/zap"
	links2 "ozon-url-shortener/app/internal/adapters/db/links"
	"ozon-url-shortener/app/internal/domain/links"
	"ozon-url-shortener/app/pkg/client/pg"
)

func (app *App) Storage() links.Storage {
	app.logger.Info("Initializing storage")

	var storage links.Storage

	if app.flags.cache {
		app.logger.Warn("Using built-in storage")
		// cache
		storage = links2.NewMemStorage()
	} else {
		// postgres
		pool, err := pg.NewPool(time.Second*3, pg.Config(app.cfg.Postgres))
		if err != nil {
			app.logger.Fatal("connect to postgres: ", zap.Error(err))
		}

		app.closers.db = pool.Close

		storage = links2.NewPgStorage(pool)
	}

	return storage
}
