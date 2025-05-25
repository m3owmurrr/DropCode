package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/m3owmurrr/dropcode/backend/internal/config"
	v1 "github.com/m3owmurrr/dropcode/backend/internal/handler"
	"github.com/m3owmurrr/dropcode/backend/internal/service"
	"github.com/m3owmurrr/dropcode/backend/pkg/broker"
	"github.com/m3owmurrr/dropcode/backend/pkg/storage"
)

func RunServer() {
	config.Load()

	broker, err := broker.NewRabbitBroker()
	if err != nil {
		log.Fatalf("can't initialize broker")
	}

	storage, err := storage.NewS3Storage()
	if err != nil {
		log.Fatalf("can't initialize storage")
	}

	service := service.NewProjectService(nil, storage, broker)
	projHandler := v1.NewProjectHandler(service)
	healthHanldler := v1.NewHealthHandler()

	router := mux.NewRouter()
	router.HandleFunc("/health", healthHanldler.CheckHealth).Methods(http.MethodGet)

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/projects/run", projHandler.RunProjectHandler).Methods(http.MethodPost)

	address := fmt.Sprintf("%s:%s", config.Cfg.Host, config.Cfg.Port)
	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Printf("Server running on %v ...", address)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server cannot run: %v", err)
		os.Exit(1)
	}
}
