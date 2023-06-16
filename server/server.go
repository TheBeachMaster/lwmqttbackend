package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	mux    *http.ServeMux
	port   string
	server *http.Server
}

func NewServer(port string, srvMux *http.ServeMux) *Server {
	srv := newHTTPServer(srvMux, port)
	return &Server{mux: srvMux, port: port, server: srv}
}

func newHTTPServer(mux *http.ServeMux, port string) *http.Server {
	s := http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &s
}

func (s *Server) Run() error {

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := s.server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	err := s.MapHTTPHandlers()
	if err != nil {
		return err
	}

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

	log.Printf("Server shutdown")

	return err
}
