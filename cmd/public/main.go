package public

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/adamkadda/ntumiwa-site/internal/handler/app"
)

func main() {

	/*
		Around here at startup, we should call our constructor
		for our App struct instead. We can then handle errors
		accordingly at this level.

		For now I've decided to keep logs simple, output to
		stdout, and just format with [PUBLIC], [ADMIN], or [API]
	*/

	templates := template.Must(template.ParseGlob("templates/*.html"))

	logger := log.New(os.Stdout, "[PUBLIC] ", log.LstdFlags)

	// complete declaring routes, do in app setup though
	router := http.NewServeMux()

	publicApp := &app.App{
		Templates: templates,
		Logger:    logger,
		Router:    router,
		DB:        nil,
	}
}
