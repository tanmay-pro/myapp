package main

import (
	"fmt"
	"log"
	"myapp/app/router"
	"myapp/config"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./docker/app/.env")
	appConf := config.AppConfig()
	appRouter := router.New()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
