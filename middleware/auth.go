package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/auth"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
)

func AuthMiddleware(client *gocloak.GoCloak) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				fmt.Println("Auth header missing")
				return echo.ErrUnauthorized
			}

			// Extract Bearer token
			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				fmt.Println("Auth token missing")
				return echo.ErrUnauthorized
			}

			// Validate token
			ctx := context.Background()

			introspect, err := client.RetrospectToken(ctx, token, auth.ClientID, auth.ClientSecret, auth.Realm)
			if err != nil || !*introspect.Active {
				log.Logger.Warn("Token invalid or introspection failed")
				return echo.ErrUnauthorized
			}

			// Get claims
			_, claims, err := client.DecodeAccessToken(ctx, token, auth.Realm)
			if err != nil {
				log.Logger.Error("Failed to decode access token", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}

			// Save user ID to context
			userID, err := claims.GetSubject()
			if err != nil {
				log.Logger.Error("Failed to get token subject", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}

			// Parse UUID
			userUUID, err := uuid.Parse(userID)
			if err != nil {
				log.Logger.Error("Failed to parse user ID", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}
			c.Set("user", userUUID)

			return next(c)
		}
	}
}
