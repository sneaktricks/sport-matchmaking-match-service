package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/auth"
	"github.com/sneaktricks/sport-matchmaking-match-service/middleware"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

type Handler struct {
	goCloakClient *gocloak.GoCloak
	oidcProvider  *oidc.Provider

	matchStore         store.MatchStore
	participationStore store.ParticipationStore
}

func New(goCloakClient *gocloak.GoCloak, oidcProvider *oidc.Provider, ms store.MatchStore, ps store.ParticipationStore) *Handler {
	return &Handler{
		goCloakClient:      goCloakClient,
		oidcProvider:       oidcProvider,
		matchStore:         ms,
		participationStore: ps,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/login", func(c echo.Context) error {
		redirectURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth?client_id=%s&response_type=code&scope=openid&redirect_uri=%s",
			auth.KeycloakURL, auth.Realm, auth.ClientID, "http://localhost:8080/callback")
		return c.Redirect(http.StatusFound, redirectURL)
	})

	g.GET("/callback", func(c echo.Context) error {
		code := c.QueryParam("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "Authorization code not found")
		}

		// Exchange the code for an access token
		token, err := exchangeCodeForToken(code)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to exchange token")
		}

		return c.String(http.StatusOK, fmt.Sprintf("Access token: %s", token))
	})

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	g.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.TimeResponse{Time: time.Now().UTC()})
	})

	// TODO: Remove this endpoint
	g.GET("/test", func(c echo.Context) error {
		username := c.QueryParam("username")
		password := c.QueryParam("password")
		jwt, err := h.goCloakClient.Login(c.Request().Context(), auth.ClientID, auth.ClientSecret, auth.Realm, username, password)
		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(http.StatusOK, jwt)
	})

	// authMiddleware := middleware.AuthMiddleware(h.goCloakClient)
	oidcMiddleware := middleware.AuthMiddlewareOIDC(
		h.oidcProvider.Verifier(auth.GetOIDCVerifierConfig()),
	)

	matchGroup := g.Group("/matches")
	matchGroup.GET("", h.FindMatches)
	matchGroup.GET("/:id", h.FindMatchByID)
	matchGroup.POST("", h.CreateMatch, oidcMiddleware)
	matchGroup.PUT("/:id", h.EditMatch, oidcMiddleware)
	matchGroup.DELETE("/:id", h.DeleteMatch, oidcMiddleware)

	matchGroup.GET("/:id/participants", h.FindParticipationsInMatch)
	matchGroup.POST("/:id/participants", h.CreateParticipation, oidcMiddleware)
	matchGroup.DELETE("/:id/participants", h.DeleteParticipation, oidcMiddleware)
}
