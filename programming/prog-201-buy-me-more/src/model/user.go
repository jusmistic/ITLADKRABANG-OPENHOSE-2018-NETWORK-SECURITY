package model

import (
	"github.com/jusmistic/programming/prog-201-buy-me-more/src/database"
)

// User struct
type User struct {
	// ID        uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Money    int64
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}

// AddUser create a new person
func AddUser(username string, password string) User {
	addUser := `INSERT INTO user (username, password, money) VALUES (?, ?, 5000)`

	database.GetDB().MustExec(addUser, username, password)

	user := User{Username: username, Password: password, Money: 5000}
	return user
}

// GetUser information
func GetUser(username string) (User, error) {
	user := User{}
	err := database.GetDB().Get(&user, "SELECT * FROM user WHERE username=? LIMIT 1", username)
	return user, err
}

// UpdateUser .
func UpdateUser(user User) error {
	updateStatment := `
	UPDATE user
	SET money = ?
	WHERE username = ?;`

	_, err := database.GetDB().Exec(updateStatment, user.Money, user.Username)
	return err
}

// Login user
func Login(username string, password string) (User, error) {
	login := "SELECT * FROM user WHERE username=? AND password=? LIMIT 1"
	user := User{}
	err := database.GetDB().Get(&user, login, username, password)
	return user, err
}

// ListUser information
func ListUser() []User {
	users := []User{}
	err := database.GetDB().Select(&users, "SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	return users
}
