package handler

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/icrowley/fake"
	"github.com/jusmistic/programming/prog-201-buy-me-more/src/model"
)

// Register page
func Register(w http.ResponseWriter, r *http.Request) {
	// session := getSession(w, r)
	// userSess := session.Values["user"]
	// if userSess != nil {
	// 	http.Redirect(w, r, "/item", 302)
	// }

	username := fake.UserName() + strconv.Itoa(rand.Intn(100000))
	password := fake.SimplePassword()

	user := model.AddUser(username, password)

	log.Printf("Username %s Password %s has registered", username, password)
	// session.Values["user"] = username
	// session.Save(r, w)
	render(w, "tmpl/register.html", user)
	// http.Redirect(w, r, "../item", 302)
}

// Login page
func Login(w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)
	switch r.Method {
	case http.MethodGet:
		username := session.Values["user"]
		log.Printf("%v", username)
		if username != nil {
			render(w, "tmpl/redirect-item.html", nil)
		}
		render(w, "tmpl/login.html", nil)
	case http.MethodPost:
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		_, err := model.Login(username, password)
		if err != nil {
			render(w, "tmpl/login.html", nil)
			return
		}
		log.Printf("Username %s Password %s has logged in", username, password)
		session.Values["user"] = username
		session.Save(r, w)
		render(w, "tmpl/redirect-item.html", nil)
	default:
		render404(w)
	}
}

// Logout user
func Logout(w http.ResponseWriter, r *http.Request) {
	destorySession(w, r)
	render(w, "tmpl/redirect-login.html", nil)
}

// ListUser is buy item
// func ListUser(w http.ResponseWriter, r *http.Request) {
// 	users := model.ListUserItem()
// 	// log.Printf("%v", users)
// 	render(w, "tmpl/users.html", users)
// }
