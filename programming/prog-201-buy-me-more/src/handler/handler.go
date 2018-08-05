package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("tN#S%DUs2Yw5wB7!vQgLS1CF@uQYoIDDc96T"))

// respondWithJSON reponse to json format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// render html page
func render(w http.ResponseWriter, filename string, data interface{}) {
	// Todo
	t, err := template.ParseFiles(filename)
	if err != nil {
		if strings.Compare(filename, "tmpl/404.html") == 0 {
			panic(err)
		}
		render404(w)
		return
	}
	t.Execute(w, data)
}

func render404(w http.ResponseWriter) {
	render(w, "tmpl/404.html", nil)
}

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, "my-app")
	if err != nil {
		log.Printf("Session error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return session
}

func destorySession(w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)

	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
