package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mskydream/audio-cloud/config"
	"github.com/mskydream/audio-cloud/db"
	"github.com/mskydream/audio-cloud/handler"
	"github.com/mskydream/audio-cloud/repository"
	"github.com/mskydream/audio-cloud/service"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Не удается загрузить config: %v\n", err)
	}
	accessTokenTTL, err := time.ParseDuration(config.Auth.AccessTokenTTL)
	if err != nil {
		log.Fatalf("Can't parse access token TTL: %s", err.Error())
	}

	refreshTokenTTL, err := time.ParseDuration(config.Auth.RefreshTokenTTL)
	if err != nil {
		log.Fatalf("Can't parse refresh token TTL: %s", err.Error())
	}

	db := db.InitDatabase(&config.Database)
	repos := repository.NewRepository(db, config.SaveDir)
	services := service.NewService(repos, accessTokenTTL, refreshTokenTTL)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	router.Run(":8080")
}
