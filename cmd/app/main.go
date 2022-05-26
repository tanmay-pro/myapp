package main

import (
	"fmt"
	"myapp/app/app"
	"myapp/app/router"
	"myapp/config"
	lr "myapp/util/logger"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./docker/app/.env")
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	application := app.New(logger)
	appRouter := router.New(application)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
