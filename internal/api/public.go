package api

import (
	"fmt"
	"net/http"

	"github.com/adamkadda/ntumiwa-site/internal/config"
	"github.com/adamkadda/ntumiwa-site/internal/models"
)

func FetchBriefBio() ([]string, error) {
	// do stuff
}

func FetchUpcomingEvents(cfg *config.Config) ([]models.Event, error) {
	url := fmt.Sprintf("http://%s:%s/public/events?type=upcoming", cfg.Host, cfg.Port)
}
