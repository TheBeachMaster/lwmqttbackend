package auth

import (
	"context"
	"log"

	models "com.thebeachmaster/mqttbackend/internal/models/auth"
)

type authRepository struct {
}

func NewAuthRepository() MQTTAuthRepository {
	return &authRepository{}
}

func (a *authRepository) Authn(ctx context.Context, authnData *models.AuthenticateDeviceInfo) error {
	log.Printf("Authenticating...\n%+v", authnData)

	return nil
}

func (a *authRepository) Authz(ctx context.Context, authzData *models.AuthorizationHTTPRequestInfo) error {
	log.Printf("Authorizing...\n%+v", authzData)

	return nil
}
