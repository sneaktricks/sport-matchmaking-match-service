package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
)

var (
	smGoogleCalendarServiceURL            = os.Getenv("GOOGLE_CALENDAR_SERVICE_URL")
	smGoogleCalendarServiceAPIKey         = os.Getenv("GOOGLE_CALENDAR_SERVICE_API_KEY")
	ErrMissingGoogleCalendarServiceURL    = fmt.Errorf("google calendar client: environment variable GOOGLE_CALENDAR_SERVICE_URL is not set")
	ErrMissingGoogleCalendarServiceAPIKey = fmt.Errorf("google calendar client: environment variable GOOGLE_CALENDAR_SERVICE_API_KEY is not set")
)

type GoogleCalendarDetails struct {
	MatchID      string         `json:"matchId"`
	MatchDetails model.MatchDTO `json:"matchDetails"`
}

type GoogleCalendarClient interface {
	CreateEventInCalendar(details *GoogleCalendarDetails) error
}

type SMGoogleCalendarClient struct{}

func NewSMGoogleCalendarClient() (client *SMGoogleCalendarClient, err error) {
	if smGoogleCalendarServiceURL == "" {
		return nil, ErrMissingGoogleCalendarServiceURL
	}
	if smGoogleCalendarServiceAPIKey == "" {
		return nil, ErrMissingGoogleCalendarServiceAPIKey
	}

	return &SMGoogleCalendarClient{}, nil
}

func (gc *SMGoogleCalendarClient) CreateEventInCalendar(details *GoogleCalendarDetails) error {
	// Convert GoogleCalendarDetails to JSON bytes
	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(details); err != nil {
		return fmt.Errorf("google calendar client: failed to encode request body: %w", err)
	}

	// Send request to Google Calendar service
	url := fmt.Sprintf("%s/calendar/events", smGoogleCalendarServiceURL)

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return fmt.Errorf("google calendar client: failed to create request: %w", err)
	}
	req.Header.Set("X-API-KEY", smGoogleCalendarServiceAPIKey)
	req.Header.Set("Content-Type", echo.MIMEApplicationJSON)

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("google calendar client: request failed: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("google calendar client: service responded with unexpected status code %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
