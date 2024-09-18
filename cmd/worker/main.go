package main

import (
	"context"
	"fmt"
	"log"

	"device-manager/internal/config"
	"device-manager/internal/container"
)

func main() {
	ctx := context.Background()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("failed to init app config: ", err)
	}

	c, err := container.InitContainer(cfg)
	if err != nil {
		log.Fatal("failed to init container: ", err)
	}

	fmt.Println("starting started")
	c.DevicesWorker.Start(ctx)
}
