package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"github.com/sneaktricks/sport-matchmaking-match-service/model/query"
)

func (h *Handler) FindMatches(c echo.Context) error {
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

	// Retrieve matches
	matches, err := h.matchStore.FindAll(
		c.Request().Context(),
		queryParams.Page,
		queryParams.Limit,
		nil,
		time.UnixMilli(0),
	)
	if err != nil {
		return HTTPError(err)
	}

	return c.JSON(http.StatusOK, matches)
}

func (h *Handler) FindMatchByID(c echo.Context) error {
	// Bind ID
	var id uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &id).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	// Find match by ID
	match, err := h.matchStore.FindByID(
		context.Background(),
		id,
	)
	if err != nil {
		return HTTPError(err)
	}

	return c.JSON(http.StatusOK, match)
}

func (h *Handler) CreateMatch(c echo.Context) error {
	// Get user ID
	userID, ok := c.Get("user").(uuid.UUID)
	if !ok {
		return HTTPError(ErrInvalidID)
	}

	createData := model.MatchCreate{}

	if err := c.Bind(&createData); err != nil {
		return HTTPError(err)
	}
	if err := c.Validate(createData); err != nil {
		return HTTPError(err)
	}

	match, err := h.matchStore.Create(
		c.Request().Context(),
		createData,
		userID,
	)
	if err != nil {
		log.Logger.Error(
			"failed to create match",
			slog.String("error", err.Error()),
		)
		return HTTPError(err)
	}

	return c.JSON(http.StatusCreated, match)
}

func (h *Handler) EditMatch(c echo.Context) error {
	// Get user ID
	// userID, ok := c.Get("user").(uuid.UUID)
	// if !ok {
	// 	return HTTPError(ErrInvalidID)
	// }
	// TODO: Check user permission to edit this resource

	// Bind ID
	var id uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &id).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	editData := model.MatchEdit{}

	if err := c.Bind(&editData); err != nil {
		return HTTPError(err)
	}
	if err := c.Validate(editData); err != nil {
		return HTTPError(err)
	}

	err = h.matchStore.Edit(
		c.Request().Context(),
		id,
		editData,
	)
	if err != nil {
		log.Logger.Warn(
			"failed to edit match",
			slog.String("error", err.Error()),
		)
		return HTTPError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteMatch(c echo.Context) error {
	// Get user ID
	// userID, ok := c.Get("user").(uuid.UUID)
	// if !ok {
	// 	return HTTPError(ErrInvalidID)
	// }
	// TODO: Check user permission to edit this resource

	// Bind ID
	var id uuid.UUID
	err := echo.PathParamsBinder(c).TextUnmarshaler("id", &id).BindError()
	if err != nil {
		return HTTPError(ErrInvalidID)
	}

	err = h.matchStore.Delete(
		c.Request().Context(),
		id,
	)
	if err != nil {
		log.Logger.Warn(
			"failed to delete match",
			slog.String("error", err.Error()),
		)
		return HTTPError(err)
	}

	return c.NoContent(http.StatusNoContent)
}
