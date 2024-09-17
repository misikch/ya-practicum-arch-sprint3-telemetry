package service

import (
	"device-manager/api"
	"device-manager/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("failed to init app config", err)
	}

	// Create service instance.
	service := &petsService{
		pets: map[int64]petstore.Pet{},
	}
	// Create generated server.
	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8081", srv); err != nil {
		log.Fatal(err)
	}
}
