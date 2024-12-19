package data

import (
	"encoding/json"
	"log"
	"net/http"

	models "com.thebeachmaster/mqttbackend/internal/models/data"
	repo "com.thebeachmaster/mqttbackend/internal/repositories/data"
)

type dataHandler struct {
	repositories repo.MQTTDataRepository
}

func NewMQTTDataHandler(_repo repo.MQTTDataRepository) MQTTDataRouteHandler {
	return &dataHandler{repositories: _repo}
}

// Default implements MQTTDataRouteHandler.
func (d *dataHandler) Default() http.HandlerFunc {
	panic("unimplemented")
}

// Store implements MQTTDataRouteHandler.
func (d *dataHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		messageData := &models.MQTTMessage{}

		err := json.NewDecoder(r.Body).Decode(messageData)
		if err != nil {
			log.Printf("Error Parsing JSON body: %s", err.Error())
		}

		if err := d.repositories.Sink(r.Context(), messageData); err != nil {
			log.Printf("Error storing MQTT message: %s", err.Error())
		}
		w.WriteHeader(http.StatusOK)
	}
}
