package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

var (
	ErrInvalidID = errors.New("invalid ID")
)

type validationErrors struct {
	Errors map[string]string `json:"errors"`
}

// Converts err into an *echo.HTTPError.
//
// If err is already of type *echo.HTTPError, the original error is returned.
//
// If err is of type validator.ValidationErrors, an *echo.HTTPError is constructed
// with status code 422 and a message that groups validation errors by field, into a
// field "errors".
//
// In case of unexpected errors, echo.NewHTTPError(http.StatusInternalServerError) is returned.
func HTTPError(err error) *echo.HTTPError {
	switch e := err.(type) {
	case *echo.HTTPError:
		return e
	case *validator.InvalidValidationError:
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	case validator.ValidationErrors:
		errs := make(map[string]string)
		for _, fe := range e {
			errs[fe.Field()] = fe.Error()
		}
		return echo.NewHTTPError(
			http.StatusUnprocessableEntity,
			validationErrors{errs},
		)
	default:
		if errors.Is(err, store.ErrMatchNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if errors.Is(err, store.ErrParticipationNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if errors.Is(err, store.ErrMatchFull) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}
		if errors.Is(err, store.ErrAlreadyParticipated) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		if errors.Is(err, ErrInvalidID) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
