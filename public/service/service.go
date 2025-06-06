package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/adamkadda/ntumiwa-site/public/pagedata"
	"github.com/adamkadda/ntumiwa-site/shared/config"
)

/*
	Thank you u/thatoneweirddev for your clear and concise
	explanation on how interfaces can be used to unit test,
	in the context of handlers and services.

	It really helped point me in a direction I feel more
	comfortable and confident to steer my project towards.

	https://www.reddit.com/r/golang/comments/1bsiojo/comment/kxfyygs
*/

type FrontendService interface {
	RenderHomePage(http.ResponseWriter, *http.Request)
}

type PublicFrontend struct {
	logger    *log.Logger
	templates *template.Template
	pages     pagedata.PageData
	client    *http.Client
}

func NewPublicFrontend(pages pagedata.PageData, client *http.Client) (*PublicFrontend, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %v", err)
	}

	logger := log.New(os.Stdout, "["+config.ServerType+"]", log.LstdFlags)

	templates := template.Must(template.ParseGlob("templates/*.html"))

	pf := &PublicFrontend{
		logger:    logger,
		templates: templates,
		pages:     pages,
		client:    client,
	}

	return pf, nil
}

func (pf *PublicFrontend) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	pf.logger.Printf("INFO: %s %s - Serving home page", r.Method, r.URL.Path)

	data, err := pf.pages.GetHomeData()
	if err != nil {
		pf.logger.Printf("ERROR: Failed to fetch home page data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = pf.templates.ExecuteTemplate(w, "home", data)
	if err != nil {
		pf.logger.Printf("ERROR: Failed to render home.html: %v", err)

		// TODO: redirect(?) to 500 error page
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
