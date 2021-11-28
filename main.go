package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=password")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to Database")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	log.Println("Pinged database!")

	// get rows from table
}