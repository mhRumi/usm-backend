package connection

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:@localhost/swesociety?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful...")
	}
}
