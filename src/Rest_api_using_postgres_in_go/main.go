package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Student struct {
	Rollno int    `json:"rollno"`
	Name   string `json:"name"`
	Course string `json:"course"`
}

func ConnectionPostgresDB() *sql.DB {
	connstring := "user=postgres password=postgrespw host=host.docker.internal port=32768 dbname=test sslmode=disable"
	// connstring := "host=localhost port=5432 user=postgres password=1234 dbname=testdb sslmode=disable"
	db, err = sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Db Connected")
	return db
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var stu []Student
	var student Student
	rows, err := db.Query("select * from student")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rollno\tName\tCourse")
		fmt.Println("----------------------")
		for rows.Next() {
			rows.Scan(&student.Rollno, &student.Name, &student.Course)
			stu = append(stu, student)
			fmt.Printf("%d\t%s\t%s\n", student.Rollno, student.Name, student.Course)
		}
	}
	json.NewEncoder(w).Encode(stu)
}

func getDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var student Student
	roll := param["rollno"]

	rows, err := db.Query("select * from student where rollno =$1", roll)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rollno\tName\tCourse")
		fmt.Println("----------------------")
		for rows.Next() {
			rows.Scan(&student.Rollno, &student.Name, &student.Course)
			fmt.Printf("%d\t%s\t%s\n", student.Rollno, student.Name, student.Course)
		}
	}
	json.NewEncoder(w).Encode(student)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	roll := param["rollno"]
	_, err := db.Exec("delete from student where rollno=$1", roll)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record deleted")
	}
}

func insertData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	_, err := db.Exec("insert into student(Rollno, name, course) values($1, $2, $3)", student.Rollno, student.Name, student.Course)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Inserted")
	}
	json.NewEncoder(w).Encode(student)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	roll := param["rollno"]
	_, err := db.Exec("update student set name=$1 where rollno=$2", student.Name, roll)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record updated")
	}

}

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/getdata", getData).Methods("GET")
	r.HandleFunc("/getdata/{rollno}", getDataById).Methods("GET")
	r.HandleFunc("/insert", insertData).Methods("POST")
	r.HandleFunc("/update/{rollno}", updateData).Methods("PUT")
	r.HandleFunc("/delete/{rollno}", deleteData).Methods("DELETE")

	fmt.Printf("Staring server at port 7000\n")
	log.Fatal(http.ListenAndServe(":7000", r))
}

func main() {

	db = ConnectionPostgresDB()
	initializeRouter()
}
