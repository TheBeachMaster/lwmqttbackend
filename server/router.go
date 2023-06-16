package server

import (
	"net/http"

	"com.thebeachmaster/mqttbackend/internal/auth/handler"
	"com.thebeachmaster/mqttbackend/internal/auth/repository"
	"com.thebeachmaster/mqttbackend/internal/auth/usecase"
)

func (s *Server) MapHTTPHandlers() error {

	s.mux.HandleFunc("/", s.MapRoutes)

	return nil
}

func (s *Server) MapRoutes(w http.ResponseWriter, r *http.Request) {

	repos := repository.NewAuthRepository()
	_usecase := usecase.NewMQTTAuthUsecase(repos)
	_handler := handler.NewAuthHTTPHandler(_usecase)

	var h http.Handler

	path := r.URL.Path

	switch path {
	case "/":
		h = _handler.Default()
	case "/authn":
		h = _handler.Authn()
	case "/authz":
		h = _handler.Authz()
	case "/sink":
		h = _handler.Store()
	default:
		h = _handler.Default()
	}

	h.ServeHTTP(w, r)
}
