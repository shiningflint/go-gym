package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "chat_dev"
	user     = "adam"
	password = ""
)

var db *sql.DB

// DbConnect will connect to the database using all constant parameters
func DbConnect() {
	var err error

	psqlInfo := fmt.Sprintf(
		"sslmode=disable host=%s port=%d dbname=%s user=%s password=%s",
		host,
		port,
		dbname,
		user,
		password,
	)
	fmt.Println("Connecting to database", psqlInfo)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db successfully!")
}
