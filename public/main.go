package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/adamkadda/ntumiwa-site/public/handler"
	"github.com/adamkadda/ntumiwa-site/public/pagedata"
	"github.com/adamkadda/ntumiwa-site/shared/api"
	"github.com/adamkadda/ntumiwa-site/shared/config"
	"github.com/adamkadda/ntumiwa-site/shared/middleware"
	"github.com/joho/godotenv"
)

//go:embed templates
var tmplDir embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	templates := template.Must(template.ParseFS(tmplDir, "templates/*.html"))

	apiClient := api.NewAPIClient(config)

	pageData, err := pagedata.New(apiClient)
	if err != nil {
		log.Fatalf("Failed to init page data: %v", err)
	}

	logger := log.New(os.Stdout, "["+config.ServerType+"]", log.LstdFlags)

	fs := http.FileServer(http.Dir("./static/"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/{$}", handler.Home(logger, templates, pageData))
	mux.Handle("/biography", handler.Biography(logger, templates, pageData))
	mux.Handle("/performances", handler.Performances(logger, templates, pageData))
	mux.Handle("/media", handler.Media(logger, templates, pageData))
	mux.Handle("/contact", handler.Contact(logger, templates, pageData))
	mux.Handle("/", handler.NotFound(logger, templates))

	stack := middleware.NewStack(
		middleware.Logging(logger),
	)

	server := http.Server{
		Addr:    config.Port,
		Handler: stack(mux),
	}

	logger.Printf("Listening on port %s\n", config.Port)
	server.ListenAndServe()
}
