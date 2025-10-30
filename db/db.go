package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=ggwp dbname=postgres sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cant ping to database:", err)
	}

	fmt.Println("Database connected succesfully")
}
