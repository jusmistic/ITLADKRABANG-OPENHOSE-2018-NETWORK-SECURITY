package main

import (
	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	app := App{}
	app.Initialize()

	app.Run(":9000")

}
