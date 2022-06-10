package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"ozon-url-shortener/app/internal/adapters/links"
)

type handler struct {
	logger  *zap.Logger
	service links.Service
}

func NewHandler(logger *zap.Logger, service links.Service) links.Handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Register(g *gin.Engine) {
	g.POST("/new", h.New)
	g.GET("/get", h.Get)

	g.GET("/:key", h.Key)
}

func (h *handler) New(c *gin.Context) {
	var json = struct {
		URL string `json:"url" binding:"required,url"`
	}{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Err{
			Err: err.Error(),
		})

		return
	}

	key, err := h.service.CreateLink(json.URL)
	if err != nil {
		h.logger.Error("new: service", zap.Error(err), zap.String("url", json.URL))

		c.JSON(http.StatusInternalServerError, internal)

		return
	}

	c.JSON(http.StatusOK, struct {
		Ok  bool   `json:"ok"`
		Key string `json:"key"`
	}{true, key})
}

func (h *handler) Get(c *gin.Context) {
	var json = struct {
		Key string `json:"key" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Err{
			Err: err.Error(),
		})

		return
	}

	link, err := h.service.GetLink(json.Key)
	if err != nil {
		h.logger.Error("get: service", zap.Error(err), zap.String("key", json.Key))

		c.JSON(http.StatusInternalServerError, internal)

		return
	}

	if link == "" {
		c.JSON(http.StatusBadRequest, Err{
			Err: "link not found",
		})
	}

	c.JSON(http.StatusOK, struct {
		Ok   bool   `json:"ok"`
		Link string `json:"link"`
	}{true, link})
}

func (h *handler) Key(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.Status(http.StatusBadRequest)

		return
	}

	link, err := h.service.GetLink(key)
	if err != nil {
		h.logger.Error("key: service", zap.Error(err), zap.String("key", key))

		c.Status(http.StatusInternalServerError)

		return
	}

	if link == "" {
		c.Status(http.StatusNotFound)

		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link)
}
