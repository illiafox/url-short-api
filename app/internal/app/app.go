package app

import (
	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/config"
)

type App struct {
	flags flags
	//
	logger *zap.Logger
	cfg    config.Config
	//
	closers struct {
		logger, db func() error
	}
}

func (app *App) Run() {
	app.Listen()
}
