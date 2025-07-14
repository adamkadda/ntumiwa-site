package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adamkadda/ntumiwa-site/api/session"
	"github.com/adamkadda/ntumiwa-site/shared/config"
	"github.com/adamkadda/ntumiwa-site/shared/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env: %v", err)
	} else {
		fmt.Println("Loaded .env successfully")
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sessionManager := session.NewSessionManager(
		session.NewInMemorySessionStore(),
		config.SessionMan.GCInterval,
		config.SessionMan.IdleExpiration,
		config.SessionMan.AbsoluteExpiration,
		config.SessionMan.CookieName,
		config.SessionMan.AdminDomain,
	)

	logger := log.New(os.Stdout, "["+config.ServerType+"]", log.LstdFlags)

	mux := http.NewServeMux()

	stack := middleware.NewStack(
		middleware.Logging(logger),
		sessionManager.Handle,
	)

	server := http.Server{
		Addr:    config.Port,
		Handler: stack(mux),
	}

	logger.Printf("Listening on port %s ...\n", config.Port)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
