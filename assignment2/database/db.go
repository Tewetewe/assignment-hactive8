package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres321"
	dbPort   = "5432"
	dbName   = "go-sql-learn"
)

var (
	db  *sql.DB
	err error
)

func StartDB() *sql.DB {
	dbCofig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err := sql.Open("postgres", dbCofig)
	if err != nil {
		log.Fatal("Error Connecting to database :", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	return db
}
