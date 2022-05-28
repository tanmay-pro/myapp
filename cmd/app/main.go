// package main

// import (
// 	"fmt"
// 	"myapp/app/app"
// 	"myapp/app/router"
// 	"myapp/config"
// 	lr "myapp/util/logger"
// 	"net/http"

// 	"time"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	godotenv.Load("./docker/app/.env")
// 	appConf := config.AppConfig()
// 	logger := lr.New(appConf.Debug)

// 	// db, err := dbConn.New(appConf)
// 	// if err != nil {
// 	// 	logger.Fatal().Err(err).Msg("")
// 	// 	return
// 	// }
// 	// if appConf.Debug {
// 	// 	db.LogMode(true)
// 	// }
// 	// application := app.New(logger, db)
// 	application := app.New(logger)
// 	appRouter := router.New(application)
// 	address := fmt.Sprintf(":%d", appConf.Server.Port)
// 	logger.Info().Msgf("Starting server %s\n", address)
// 	s := &http.Server{
// 		Addr:         address,
// 		Handler:      appRouter,
// 		ReadTimeout:  30 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 	}
// 	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 		logger.Fatal().Err(err).Msg("Server startup failed")
// 	}
// }

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"myapp/app/app"
	"myapp/app/router"
	"myapp/config"
	lr "myapp/util/logger"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./docker/app/.env")
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)

	// db, err := dbConn.New(appConf)
	// if err != nil {
	// 	logger.Fatal().Err(err).Msg("")
	// 	return
	// }
	// if appConf.Debug {
	// 	db.LogMode(true)
	// }
	// application := app.New(logger, db)

	application := app.New(logger)
	appRouter := router.New(application)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %s\n", address)
	// s := &http.Server{
	// 	Addr:         address,
	// 	Handler:      appRouter,
	// 	ReadTimeout:  30 * time.Second,
	// 	WriteTimeout: 30 * time.Second,
	// 	IdleTimeout:  120 * time.Second,
	// }

	// Implementting https

	addr := flag.String("addr", address, "HTTPS network address")
	certFile := flag.String("certfile", "server.crt", "certificate PEM file")
	keyFile := flag.String("keyfile", "server.key", "key PEM file")
	flag.Parse()

	s := &http.Server{
		Addr:    *addr,
		Handler: appRouter,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	if err := s.ListenAndServeTLS(*certFile, *keyFile); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
	// if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	logger.Fatal().Err(err).Msg("Server startup failed")
	// }
}
