package handlers

import (
	"net/http"

	"github.com/sanyathecreator/godern-web/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	const fn = "home.page.tmpl"
	render.RenderTemplate(w, fn)
}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {
	const fn = "about.page.tmpl"
	render.RenderTemplate(w, fn)
}
