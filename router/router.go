package router

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sneaktricks/sport-matchmaking-match-service/middleware"
)

func New(goCloakClient *gocloak.GoCloak) *echo.Echo {
	e := echo.New()
	e.Pre(echomiddleware.RemoveTrailingSlash())
	e.Use(echomiddleware.BodyLimit("2M"))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORS()) // TODO: Configure CORS for production
	e.Validator = NewValidator()

	e.Use(middleware.AuthMiddleware(goCloakClient))

	return e
}
