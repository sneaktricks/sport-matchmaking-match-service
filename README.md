# Sport Matchmaking Match Service

The Match Service API for Sport Matchmaking manages matches and participations therein.

## Required Environment Variables

This microservice depends on Keycloak, a PostgreSQL database, and the Sport Matchmaking Notification Service to function.
Therefore, the following environment variables must be defined:

### Keycloak

- KEYCLOAK_URL
    - The URL where Keycloak can be accessed, e.g. https://keycloak:8080
- KEYCLOAK_REALM
    - Keycloak realm that the application uses, e.g. sport-matchmaking
- KEYCLOAK_CLIENT_ID
    - Keycloak client ID of this microservice, e.g. match-service

### PostgreSQL

- PGUSER
    - PostgreSQL user, e.g. user
- PGPASSWORD
    - PostgreSQL password, e.g. password
- PGHOST
    - PostgreSQL host, e.g. sport-matchmaking-postgresql-db
- PGPORT
    - PostgreSQL port, e.g. 5432
- PGDATABASE
    - PostgreSQL database, e.g. database

### Notification Service

- NOTIFICATION_SERVICE_URL
    - The URL where Notification Service can be accessed, e.g. http://sport-matchmaking-notification-service-service:8080
- NOTIFICATION_SERVICE_API_KEY
    - API key used to authenticate with notification service

## Running on Kubernetes

To run Match Service on Kubernetes, the following secrets must be present in the cluster:

- match-service-keycloak-credentials
    - url
    - realm
    - clientId
- postgres-credentials
    - user
    - password
    - host
    - port
    - database
- notification-service-match-service-secret
    - apiKey

With the secrets in place, execute `kubectl apply -f ./kubernetes/` to start the service.