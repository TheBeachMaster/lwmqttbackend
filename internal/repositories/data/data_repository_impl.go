package data

import (
	"context"
	"log"

	"com.thebeachmaster/mqttbackend/internal/models/data"
)

type dataRepository struct {
}

func NewDataRepository() MQTTDataRepository {
	return &dataRepository{}
}

// Default implements MQTTDataRepository.
func (d *dataRepository) Default(ctx context.Context) error {
	log.Printf("Default Data")

	return nil
}

// Sink implements MQTTDataRepository.
func (d *dataRepository) Sink(ctx context.Context, messageData *data.MQTTMessage) error {
	log.Printf("Storing...\n%+v", messageData)

	return nil
}
