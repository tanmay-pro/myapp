package main

import (
	"fmt"
	"log"
	"myapp/config"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./docker/app/.env")
	appConf := config.AppConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
