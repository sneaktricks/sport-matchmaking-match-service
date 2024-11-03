package middleware

import (
	"context"
	"log/slog"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/auth"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
)

func AuthMiddleware(client *gocloak.GoCloak) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			// Extract Bearer token
			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				return echo.ErrUnauthorized
			}

			// Validate token
			ctx := context.Background()
			introspect, err := client.RetrospectToken(ctx, token, auth.ClientID, auth.ClientSecret, auth.Realm)
			if err != nil || introspect.Active == nil || !*introspect.Active {
				log.Logger.Warn("Token invalid or introspection failed", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}

			// Get user info
			userInfo, err := client.GetUserInfo(ctx, token, auth.Realm)
			if err != nil {
				log.Logger.Error("Failed to retrieve user info", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}

			// Save user info to context
			c.Set("user", userInfo)

			return next(c)
		}
	}
}
