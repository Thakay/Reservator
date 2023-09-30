package handlers

import (
	"github.com/thakay/Reservator/pkg/config"
	"github.com/thakay/Reservator/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
