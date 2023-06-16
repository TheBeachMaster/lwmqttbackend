package auth

import (
	"context"

	"com.thebeachmaster/mqttbackend/internal/auth/models"
)

type AuthUseCase interface {
	Sink(ctx context.Context, messageData *models.MQTTMessage) error
	Authn(ctx context.Context, authnData *models.AuthenticateDeviceInfo) error
	Authz(ctx context.Context, authzData *models.AuthorizationHTTPRequestInfo) error
	Default(ctx context.Context) error
}
