package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID     int    `json:"id"`
	Name   string `json:"name"`
	course string `json:"course"`
}

func ConnectionPostgresDB() *sql.DB {
	connstring := "host=localhost port=5432 user=postgres password=1234 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func getDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

}

func insertData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/getdata", getData).Methods("GET")
	r.HandleFunc("/getdata/{id}", getDataById).Methods("GET")
	r.HandleFunc("/insert", insertData).Methods("POST")
	r.HandleFunc("/update/{id}", updateData).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteData).Methods("DELETE")

	fmt.Printf("Staring server at port 8080\n")
	log.Fatal(http.ListenAndServe(":7000", r))
}

func main() {
	initializeRouter()
}
