package usecase

import (
	"context"

	"com.thebeachmaster/mqttbackend/internal/auth"
	"com.thebeachmaster/mqttbackend/internal/auth/models"
)

type authUsecase struct {
	repo auth.AuthRepository
}

func NewMQTTAuthUsecase(r auth.AuthRepository) auth.AuthUseCase {
	return &authUsecase{repo: r}
}

func (a *authUsecase) Sink(ctx context.Context, messageData *models.MQTTMessage) error {
	return a.repo.Sink(ctx, messageData)
}

func (a *authUsecase) Authn(ctx context.Context, authnData *models.AuthenticateDeviceInfo) error {
	return a.repo.Authn(ctx, authnData)
}

func (a *authUsecase) Authz(ctx context.Context, authzData *models.AuthorizationHTTPRequestInfo) error {
	return a.repo.Authz(ctx, authzData)
}
func (a *authUsecase) Default(ctx context.Context) error {
	return a.repo.Default(ctx)
}
