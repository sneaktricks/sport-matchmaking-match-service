package handler

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
	"github.com/sneaktricks/sport-matchmaking-match-service/model/query"
)

func (h *Handler) FindParticipationsInMatch(c echo.Context) error {
	// Bind ID
	var id uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &id).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	// Bind and validate query params
	queryParams := query.PaginationParams{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &queryParams); err != nil {
		return HTTPError(err)
	}
	if err := c.Validate(queryParams); err != nil {
		return HTTPError(err)
	}
	// Set default param values if not defined
	if queryParams.Limit == 0 {
		queryParams.Limit = 25
	}
	if queryParams.Page == 0 {
		queryParams.Page = 1
	}

	// Retrieve participations
	participations, err := h.participationStore.FindAllInMatch(
		c.Request().Context(),
		id,
		queryParams.Page,
		queryParams.Limit,
	)
	if err != nil {
		return HTTPError(err)
	}

	return c.JSON(http.StatusOK, participations)
}

func (h *Handler) CreateParticipation(c echo.Context) error {
	// Get user ID
	userID, ok := c.Get("user").(string)
	if !ok {
		return HTTPError(ErrInvalidID)
	}

	// Bind ID
	var matchID uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &matchID).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	participation, err := h.participationStore.Create(
		c.Request().Context(),
		matchID,
		userID,
	)
	if err != nil {
		log.Logger.Error(
			"failed to participate",
			slog.String("error", err.Error()),
		)
		return HTTPError(err)
	}

	return c.JSON(http.StatusCreated, participation)
}

func (h *Handler) DeleteParticipation(c echo.Context) error {
	// Get user ID
	userID, ok := c.Get("user").(string)
	if !ok {
		return HTTPError(ErrInvalidID)
	}

	// Bind matchID
	var matchID uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &matchID).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	err = h.participationStore.Delete(
		c.Request().Context(),
		matchID,
		userID,
	)
	if err != nil {
		log.Logger.Warn(
			"failed to unparticipate",
			slog.String("error", err.Error()),
		)
		return HTTPError(err)
	}

	return c.NoContent(http.StatusNoContent)
}
