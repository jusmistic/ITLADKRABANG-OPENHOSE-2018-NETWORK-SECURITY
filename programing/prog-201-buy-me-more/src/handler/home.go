package handler

import (
	"net/http"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "tmpl/home.html", nil)
}
