package pagedata

import "github.com/adamkadda/ntumiwa-site/shared/models"

type HomeData struct {
	BriefBio       []string       `json:"briefBio"`
	UpcomingEvents []models.Event `json:"events"`
}
