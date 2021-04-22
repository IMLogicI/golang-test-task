package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Service struct {
	Router *mux.Router
	Server http.Server
}

func New() *Service {
	router := mux.NewRouter()
	router.HandleFunc("/", UserIdHandler()).Methods(http.MethodPost)

	srv := Service{
		Router: router,
	}

	srv.Server = http.Server{
		Addr:    ":8080",
		Handler: srv.Router,
	}

	return &srv
}

func (s *Service) Start() {

	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func (s *Service) Stop(context context.Context) {
	if err := s.Server.Shutdown(context); err != nil {
		log.Printf("error when server is shutting down: %s", err)
		return
	}
}
