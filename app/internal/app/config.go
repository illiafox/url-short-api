package app

import (
	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/config"
)

func (app *App) ReadConfig() {
	app.logger.Info("Reading config")

	cfg, err := config.New(app.flags.config)
	if err != nil {
		app.logger.Fatal("read config",
			zap.String("path", app.flags.config),
			zap.Error(err),
		)
	}

	app.cfg = cfg
}
