package auth

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

var (
	KeycloakURL  = os.Getenv("KEYCLOAK_URL")
	Realm        = os.Getenv("KEYCLOAK_REALM")
	ClientID     = os.Getenv("KEYCLOAK_CLIENT_ID")
	ClientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
)

func NewGoCloakClient() *gocloak.GoCloak {
	return gocloak.NewClient(KeycloakURL)
}
