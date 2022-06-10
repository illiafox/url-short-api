package app

import (
	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/adapters/links"
	"ozon-url-shortener/app/internal/config"
)

type App struct {
	flags flags
	//
	logger *zap.Logger
	cfg    config.Config
	//
	service links.Service
	//
	closers struct {
		logger  func()
		logfile func()
		db      func()
	}
}

func (app *App) Run() {
	app.Logger()
	app.ReadConfig()

	app.Service()
	app.Listen()
}
