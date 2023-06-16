package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"com.thebeachmaster/mqttbackend/internal/auth"
	"com.thebeachmaster/mqttbackend/internal/auth/models"
)

type authHandler struct {
	usecase auth.AuthUseCase
}

func NewAuthHTTPHandler(u auth.AuthUseCase) auth.MQTTRouteHandlers {
	return &authHandler{usecase: u}
}

func (a *authHandler) Authn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		authInfo := &models.AuthenticateDeviceInfo{}

		authNResponse := &models.AuthNResponse{}
		authNResponse.IsSuperUser = false
		authNResponse.Result = "allow"

		err := json.NewDecoder(r.Body).Decode(authInfo)
		if err != nil {
			log.Printf("Error Parsing JSON body: %s", err.Error())
			authNResponse.IsSuperUser = false
			authNResponse.Result = "deny"
		}
		if err = a.usecase.Authn(r.Context(), authInfo); err != nil {
			log.Printf("Authentication failed")
			authNResponse.IsSuperUser = false
			authNResponse.Result = "deny"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authNResponse)
	}
}

func (a *authHandler) Authz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		authInfo := &models.AuthorizationHTTPRequestInfo{}

		authZResponse := &models.AuthZResponse{}
		authZResponse.Result = "allow"

		err := json.NewDecoder(r.Body).Decode(authInfo)
		if err != nil {
			log.Printf("Error Parsing JSON body: %s", err.Error())
			authZResponse.Result = "deny"
		}
		if err = a.usecase.Authz(r.Context(), authInfo); err != nil {
			log.Printf("Authorization failed")
			authZResponse.Result = "deny"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authZResponse)
	}
}

func (a *authHandler) Store() http.HandlerFunc {
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

		if err := a.usecase.Sink(r.Context(), messageData); err != nil {
			log.Printf("Error storing MQTT message: %s", err.Error())
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (a *authHandler) Default() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := &models.HealthStatus{
			Status: "imok",
			Time:   time.Now().String(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(status)
	}
}
