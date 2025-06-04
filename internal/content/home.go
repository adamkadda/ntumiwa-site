package content

import "github.com/adamkadda/ntumiwa-site/internal/models"

type HomeData struct {
	BriefBio       []string
	UpcomingEvents []models.Event
}
