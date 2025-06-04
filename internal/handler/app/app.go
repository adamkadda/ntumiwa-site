package app

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/adamkadda/ntumiwa-site/internal/config"
	"github.com/adamkadda/ntumiwa-site/internal/content"
)

type App struct {
	Templates  *template.Template
	Logger     *log.Logger
	DB         *sql.DB
	Router     *http.ServeMux
	Pages      *content.PageDataCache
	Config     *config.Config
	HTTPClient *http.Client
}

func NewApp() *App {
	// prepare templates

	// prepare logger

	// prepare DB (handle expected nil case?)

	// init mux/router, and routes

	// init page data cache

	// init config
}
