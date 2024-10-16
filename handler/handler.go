package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

type Handler struct {
	matchStore store.MatchStore
}

func New(ms store.MatchStore) *Handler {
	return &Handler{
		matchStore: ms,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now().UTC()})
	})

	matchGroup := g.Group("/matches")
	matchGroup.GET("", h.FindMatches)
	matchGroup.GET("/:id", h.FindMatchByID)
	matchGroup.POST("", h.CreateMatch)
	matchGroup.PUT("/:id", h.EditMatch)
	matchGroup.DELETE("/:id", h.DeleteMatch)
}
