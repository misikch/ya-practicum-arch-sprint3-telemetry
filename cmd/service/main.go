package main

import (
	"fmt"
	"log"
	"net/http"

	"device-manager/api"
	"device-manager/internal/config"
	"device-manager/internal/container"
	"device-manager/internal/handler"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("failed to init app config: ", err)
	}

	c, err := container.InitContainer(cfg)
	if err != nil {
		log.Fatal("failed to init container: ", err)
	}

	// Create handlers
	handlers := handler.NewHandler(c.TelemetryService, c.Logger)

	// Create generated server.
	srv, err := api.NewServer(handlers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("starting service at :%s ...\n", cfg.ServicePort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServicePort), srv); err != nil {
		log.Fatal(err)
	}
}
