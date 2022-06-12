package app

import (
	"log"
	"os"

	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/config"
)

func (app *App) ReadConfig() {
	// // config
	cfg, err := config.New(app.flags.config)
	if err != nil {
		app.logger.Error("read config",
			zap.String("path", app.flags.config),
			zap.Error(err),
		)

		// close logger
		err = app.closers.logger()
		if err != nil {
			log.Fatalf("close logger: %v", err)
		}

		os.Exit(1)
	}

	app.cfg = cfg
}
