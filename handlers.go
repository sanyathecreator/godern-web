package main

import (
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	const fn = "home.page.tmpl"
	renderTemplate(w, fn)
}

// About is the home page handler
func About(w http.ResponseWriter, r *http.Request) {
	const fn = "about.page.tmpl"
	renderTemplate(w, fn)
}
