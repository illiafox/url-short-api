package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ozon-url-shortener/app/internal/adapters/links/api"
)

func (app *App) Handler() http.Handler {
	h := api.NewHandler(app.logger.Named("api"), app.service)

	if !app.flags.debug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(gin.Recovery())

	if app.flags.debug {
		g.Use(gin.Logger())
	}

	h.Register(g)

	return g
}
