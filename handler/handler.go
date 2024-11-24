package handler

import (
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/auth"
	"github.com/sneaktricks/sport-matchmaking-match-service/integrations/notification"
	"github.com/sneaktricks/sport-matchmaking-match-service/middleware"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

type Handler struct {
	oidcProvider       *oidc.Provider
	notificationClient notification.NotificationClient

	matchStore         store.MatchStore
	participationStore store.ParticipationStore
}

func New(oidcProvider *oidc.Provider, notificationClient notification.NotificationClient, ms store.MatchStore, ps store.ParticipationStore) *Handler {
	return &Handler{
		oidcProvider:       oidcProvider,
		notificationClient: notificationClient,
		matchStore:         ms,
		participationStore: ps,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sport Matchmaking Match Service API")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now().UTC()})
	})

	authMiddleware := middleware.AuthMiddleware(
		h.oidcProvider.Verifier(auth.GetOIDCVerifierConfig()),
	)

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
