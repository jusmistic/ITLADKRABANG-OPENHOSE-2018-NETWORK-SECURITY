package handler

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/icrowley/fake"

	"github.com/jusmistic/programming/prog-201-buy-me-more/src/model"
)

// ViewItem user
func ViewItem(w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)
	userSess := session.Values["user"]
	if userSess == nil {
		Logout(w, r)
		return
	}
	user, err := model.GetUser(userSess.(string))
	if err != nil {
		Logout(w, r)
		return
		// panic(err)
	}

	userItems := model.GetUserItem(user.Username)

	render(w, "tmpl/user_item.html", userItems)
}

// BuyItem buy
func BuyItem(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	session := getSession(w, r)
	if session.Values["user"] == nil {
		render(w, "tmpl/404.html", nil)
		return
	}
	user, err := model.GetUser(session.Values["user"].(string))
	if err != nil {
		panic(err)
	}

	if user.Money <= 0 {
		usageTime := time.Now().Sub(startTime)

		render(w, "tmpl/no_money.html", usageTime)
		return
	}

	// Simulate long operation delay 10-20ms
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)+10))
	// Add item to user !!!
	model.CreateUserItem(user.Username, fake.ProductName())

	user.Money = user.Money - 2500
	model.UpdateUser(user)
	usageTime := time.Now().Sub(startTime)
	log.Printf("Username: %s buy a new item.", user.Username)
	render(w, "tmpl/buy_success.html", usageTime)
}
