package auth

import (
	"context"

	models "com.thebeachmaster/mqttbackend/internal/models/auth"
)

type MQTTAuthRepository interface {
	Authn(ctx context.Context, authnData *models.AuthenticateDeviceInfo) error
	Authz(ctx context.Context, authzData *models.AuthorizationHTTPRequestInfo) error
}
