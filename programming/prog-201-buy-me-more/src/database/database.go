package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// Connect create a database connection
func Connect(uri string) {
	var err error

	db, err = sqlx.Connect("sqlite3", "file:test.db")
	if err != nil {
		log.Fatal("connect database error: ", err)
		panic("failed to connect database")
	}

	log.Println("Database is connected !")

	prepare()
}

// Close close a database connection
func Close() {
	var err error
	err = db.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func prepare() {
	log.Println("Prepare database")
	schema := `CREATE TABLE IF NOT EXISTS user (
		username varchar(64) DEFAULT NULL,
		password varchar(128) DEFAULT NULL,
		money int(11) DEFAULT NULL
	  );

	  CREATE TABLE IF NOT EXISTS user_item (
		username varchar(64) DEFAULT NULL,
		item varchar(64) DEFAULT NULL
	  );`
	db.MustExec(schema)
}

// GetDB get database
func GetDB() *sqlx.DB {
	return db
}
