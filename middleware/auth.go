package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
)

func AuthMiddleware(verifier *oidc.IDTokenVerifier) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				fmt.Println("Authorization header missing")
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

			idToken, err := verifier.Verify(ctx, token)
			if err != nil {
				log.Logger.Warn("Invalid token", slog.String("error", err.Error()))
				return echo.ErrUnauthorized
			}

			// Save user ID to context
			userID := idToken.Subject

			// Check if user ID is empty
			if userID == "" {
				log.Logger.Error("Encountered invalid user ID in subject")
				return echo.ErrUnauthorized
			}
			c.Set("user", userID)

			return next(c)
		}
	}
}
