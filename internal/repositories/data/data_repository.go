package data

import (
	"context"

	models "com.thebeachmaster/mqttbackend/internal/models/data"
)

type MQTTDataRepository interface {
	Sink(ctx context.Context, messageData *models.MQTTMessage) error
	Default(ctx context.Context) error
}
