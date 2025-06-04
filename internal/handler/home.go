package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adamkadda/ntumiwa-site/internal/handler/app"
)

func Home(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Printf("INFO: %s %s - Serving home page", r.Method, r.URL.Path)

		data := app.Pages.Home.Get()
		if data == nil {
			/*
				Later make sure to use a .env file instead of hardcoding.
				Add it to .gitignore to ensure it doesn't appear in the repo, ever.

				Then, at startup load it in using "github.com/joho/godotenv/autoload",
				into a Config, then into the App.

				That last bit however is still up for discussion.
			*/

			apiURL := "lol"

			// Attempt to ping endpoint, use app.Config, handle (potential) error
			resp, err := http.Get(apiURL)
			if err != nil {
				app.Logger.Printf("ERROR: Failed to fetch home page data: %v", err)
				http.Error(w, "Could not fetch home page data", http.StatusInternalServerError)
			}
			defer resp.Body.Close()

			// Handle unexpected status codes, don't propagate yet
			if resp.StatusCode != http.StatusOK {
				app.Logger.Printf("ERROR: Unexpected status received: %v", resp.Status)
				http.Error(w, "Backend API error", http.StatusBadGateway)
			}
		}

		/*
			Apparently, it is by design in Go that we call the ExecuteTemplate()
			method per request. Calling the method is designed to be cheap and fast,
			so unless we really need to cache the result, this approach should be OK.
		*/

		err := app.Templates.ExecuteTemplate(w, "home", data)
		if err != nil {
			app.Logger.Printf("ERROR: Failed to render home.html: %v", err)

			// decide if we want a dedicated 500 error page
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
