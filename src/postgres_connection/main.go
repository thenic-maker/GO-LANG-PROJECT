package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//conninfo := "user=postgres password=postgrespw host=host.docker.internal port=32768 sslmode=disable"
	conninfo := "postgres://postgres:postgrespw@localhost:32768?sslmode=disable"
	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		log.Fatal(err)
	}
	dbName := "testd"
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		//handle the error
		log.Fatal(err)
	}

	//Then execute your query for creating table
	_, err = db.Exec("CREATE TABLE example ( id integer, username varchar(255) )")

	if err != nil {
		log.Fatal(err)
	}

}
