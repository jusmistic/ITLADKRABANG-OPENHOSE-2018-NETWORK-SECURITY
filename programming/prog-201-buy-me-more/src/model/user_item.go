package model

import (
	"github.com/jusmistic/programming/prog-201-buy-me-more/src/database"
)

// UserItem struct
type UserItem struct {
	Item     string
	Username string
	User     `db:"user"`
}

// CreateUserItem user item
func CreateUserItem(username string, itemName string) {
	addItem := `INSERT INTO user_item (username, item) VALUES (?, ?)`

	database.GetDB().MustExec(addItem, username, itemName)
}

// GetUserItem user item
func GetUserItem(username string) []UserItem {
	userItem := []UserItem{}
	database.GetDB().Select(&userItem, `
	SELECT user_item.*,
	user.username "user.username",
	user.password "user.password",
	user.money "user.money"
	FROM user_item 
	JOIN user ON user_item.username = user.username
	WHERE user_item.username = ?`, username)
	return userItem
}

func ListUserItem() []UserItem {
	userItems := []UserItem{}
	database.GetDB().Select(&userItems, `
	SELECT user_item.*,
	user.username "user.username",
	user.password "user.password",
	user.money "user.money"
	FROM user_item 
	JOIN user ON user_item.username = user.username`)
	return userItems
}
