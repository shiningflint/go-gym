package connection

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

// DbConnect will connect to the database using all constant parameters
func DbConnect() {
	psqlInfo := fmt.Sprintf(
		"sslmode=disable host=%s port=%d dbname=%s user=%s password=%s",
		host,
		port,
		dbname,
		user,
		password,
	)
	fmt.Println("Connecting to database", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db successfully!")
}
