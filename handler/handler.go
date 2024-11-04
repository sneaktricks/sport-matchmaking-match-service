package handler

import (
	"net/http"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/middleware"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

type Handler struct {
	goCloakClient *gocloak.GoCloak

	matchStore         store.MatchStore
	participationStore store.ParticipationStore
}

func New(goCloakClient *gocloak.GoCloak, ms store.MatchStore, ps store.ParticipationStore) *Handler {
	return &Handler{
		goCloakClient:      goCloakClient,
		matchStore:         ms,
		participationStore: ps,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now().UTC()})
	})

	authMiddleware := middleware.AuthMiddleware(h.goCloakClient)

	matchGroup := g.Group("/matches")
	matchGroup.GET("", h.FindMatches)
	matchGroup.GET("/:id", h.FindMatchByID)
	matchGroup.POST("", h.CreateMatch, authMiddleware)
	matchGroup.PUT("/:id", h.EditMatch, authMiddleware)
	matchGroup.DELETE("/:id", h.DeleteMatch, authMiddleware)

	matchGroup.GET("/:id/participants", h.FindParticipationsInMatch)
	matchGroup.POST("/:id/participants", h.CreateParticipation, authMiddleware)
	matchGroup.DELETE("/:id/participants", h.DeleteParticipation, authMiddleware)
}
