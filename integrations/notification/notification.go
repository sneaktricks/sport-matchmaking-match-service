package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
)

var (
	smNotificationServiceURL           = os.Getenv("NOTIFICATION_SERVICE_URL")
	ErrMissingSMNotificationServiceURL = errors.New("SM notification client: enviroment variable NOTIFICATION_SERVICE_URL is not set")
)

type NotificationDetails struct {
	UserIDs      []string       `json:"userIds"`
	MatchDetails model.MatchDTO `json:"matchDetails"`
}

type NotificationClient interface {
	NotifyUsersAboutMatchUpdate(details *NotificationDetails) error
}

type SMNotificationClient struct {
	url string
}

func NewSMNotificationClient() (client *SMNotificationClient, err error) {
	if smNotificationServiceURL == "" {
		return nil, ErrMissingSMNotificationServiceURL
	}

	return &SMNotificationClient{url: smNotificationServiceURL}, nil
}

func (nc *SMNotificationClient) NotifyUsersAboutMatchUpdate(details *NotificationDetails) error {
	// Convert NotificationDetails to JSON bytes
	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(details); err != nil {
		return fmt.Errorf("SM notification client: failed to send notification: %w", err)
	}

	// Send notification request
	url := fmt.Sprintf("%s/notify", smNotificationServiceURL)
	resp, err := http.Post(url, echo.MIMEApplicationJSON, reqBody)
	if err != nil {
		return fmt.Errorf("SM notification client: notification request failed: %w", err)
	}

	// Check status code
	if resp.StatusCode < 200 || 299 < resp.StatusCode {
		return fmt.Errorf("SM notification client: service responded with unexpected status code %d", resp.StatusCode)
	}

	return nil
}
