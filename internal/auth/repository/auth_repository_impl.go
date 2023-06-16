package repository

import (
	"context"
	"log"

	"com.thebeachmaster/mqttbackend/internal/auth"
	"com.thebeachmaster/mqttbackend/internal/auth/models"
)

type authRepository struct {
}

func NewAuthRepository() auth.AuthRepository {
	return &authRepository{}
}

func (a *authRepository) Sink(ctx context.Context, messageData *models.MQTTMessage) error {
	log.Printf("Storing...\n%+v", messageData)

	return nil
}

func (a *authRepository) Authn(ctx context.Context, authnData *models.AuthenticateDeviceInfo) error {
	log.Printf("Authenticating...\n%+v", authnData)

	return nil
}

func (a *authRepository) Authz(ctx context.Context, authzData *models.AuthorizationHTTPRequestInfo) error {
	log.Printf("Authorizing...\n%+v", authzData)

	return nil
}

func (a *authRepository) Default(ctx context.Context) error {
	log.Printf("Default Data")

	return nil
}
