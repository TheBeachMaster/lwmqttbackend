package auth

import (
	"encoding/json"
	"log"
	"net/http"

	models "com.thebeachmaster/mqttbackend/internal/models/auth"
	repo "com.thebeachmaster/mqttbackend/internal/repositories/auth"
)

type authHandler struct {
	repository repo.MQTTAuthRepository
}

func NewAuthHTTPHandler(_repo repo.MQTTAuthRepository) MQTTAuthRouteHandlers {
	return &authHandler{repository: _repo}
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
		if err = a.repository.Authn(r.Context(), authInfo); err != nil {
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
		if err = a.repository.Authz(r.Context(), authInfo); err != nil {
			log.Printf("Authorization failed")
			authZResponse.Result = "deny"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authZResponse)
	}
}
