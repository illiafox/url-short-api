package app

import (
	"ozon-url-shortener/app/internal/domain/links"
)

func (app *App) Service() {
	app.logger.Info("Initializing service")

	app.service = links.NewService(app.Storage())
}
