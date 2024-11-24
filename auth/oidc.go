package auth

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
)

var (
	KeycloakURL = os.Getenv("KEYCLOAK_URL")
	Realm       = os.Getenv("KEYCLOAK_REALM")
	ClientID    = os.Getenv("KEYCLOAK_CLIENT_ID")
)

func checkEnv() error {
	missingEnvs := make([]string, 0)
	if KeycloakURL == "" {
		missingEnvs = append(missingEnvs, "KEYCLOAK_URL")
	}
	if Realm == "" {
		missingEnvs = append(missingEnvs, "KEYCLOAK_REALM")
	}
	if ClientID == "" {
		missingEnvs = append(missingEnvs, "KEYCLOAK_CLIENT_ID")
	}

	if len(missingEnvs) > 0 {
		return fmt.Errorf("the following Keycloak environment variables are undefined: %s", strings.Join(missingEnvs, ", "))
	}

	return nil
}

func NewOIDCProvider() (provider *oidc.Provider, err error) {
	if err := checkEnv(); err != nil {
		log.Logger.Error(err.Error())
	}

	provider, err = oidc.NewProvider(
		context.Background(),
		fmt.Sprintf("%s/realms/%s", KeycloakURL, Realm),
	)
	return provider, err
}

func GetOIDCVerifierConfig() *oidc.Config {
	return &oidc.Config{
		ClientID: ClientID,
	}
}
