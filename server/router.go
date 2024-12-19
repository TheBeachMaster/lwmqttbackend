package server

import (
	"net/http"

	authHandler "com.thebeachmaster/mqttbackend/internal/handlers/auth"
	dataHandler "com.thebeachmaster/mqttbackend/internal/handlers/data"
	authRepo "com.thebeachmaster/mqttbackend/internal/repositories/auth"
	dataRepo "com.thebeachmaster/mqttbackend/internal/repositories/data"
)

func (s *Server) MapHTTPHandlers() error {

	s.mux.HandleFunc("/", s.MapRoutes)

	return nil
}

func (s *Server) MapRoutes(w http.ResponseWriter, r *http.Request) {

	_authRepo := authRepo.NewAuthRepository()
	_dataRepo := dataRepo.NewDataRepository()

	_authHandler := authHandler.NewAuthHTTPHandler(_authRepo)
	_dataHandler := dataHandler.NewMQTTDataHandler(_dataRepo)

	var h http.Handler

	path := r.URL.Path

	switch path {
	case "/":
		h = _dataHandler.Default()
	case "/authn":
		h = _authHandler.Authn()
	case "/authz":
		h = _authHandler.Authz()
	case "/sink":
		h = _dataHandler.Store()
	default:
		h = _dataHandler.Default()
	}

	h.ServeHTTP(w, r)
}
