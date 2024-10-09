package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now()})
	})
}
