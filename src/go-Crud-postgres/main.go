package main

import (
	"bufio"
	"database/sql"
	"fmt"

	// "log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func ConnectionPostgresDB() *sql.DB {
	connstring := "host=localhost port=5432 user=postgres password=1234 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// func createDb(db *sql.DB) {
// 	dbName := "testdb"
// 	_, err := db.Exec("create database " + dbName)
// 	if err != nil {
// 		//handle the error
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Database Created")
// 	}
// }

func insertIntoDB(db *sql.DB, rollno int, name string, course string) {
	_, err := db.Exec("insert into student(Rollno, name, course) values($1, $2, $3)", rollno, name, course)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Inserted")
	}
}

func updateData(db *sql.DB, rollno int, name string) {
	_, err := db.Exec("update student set name=$1 where rollno=$2", name, rollno)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record updated")
	}
}

func deleteData(db *sql.DB, rollno int) {
	_, err := db.Exec("delete from student where rollno=$1", rollno)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record deleted")
	}
}

func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}

func main() {
	var choice string
	var rollno int
	var name string
	var course string
	db := ConnectionPostgresDB()
	dbName := "testdb"
	_, err := db.Exec("create database " + dbName)
	if err != nil {
		//handle the error
		fmt.Println("hello")
		fmt.Println(err)
	}
	_, err = db.Exec("\\c " + dbName)
	if err != nil {
		//handle the error
		fmt.Println(err)
	}
	//Then execute your query for creating table
	_, err = db.Exec("CREATE TABLE student( id integer, Name TEXT, Course TEXT )")

	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("\nEnter the choice")

		fmt.Println("1. Insert data in Postgres DB")
		fmt.Println("2. Read data from Postgres DB")
		fmt.Println("3. Update data in Postgres DB")
		fmt.Println("4. Delete data from Postgres DB")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			fmt.Println("Enter Rollno")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			fmt.Println("Enter name")
			name = getInput()
			fmt.Println("Enter course")
			course = getInput()
			insertIntoDB(db, rollno, name, course)
		case "2":
			rows, err := db.Query("select * from student")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rollno\tName\tCourse")
				fmt.Println("----------------------")
				for rows.Next() {
					rows.Scan(&rollno, &name, &course)
					fmt.Printf("%d\t%s\t%s\n", rollno, name, course)
				}
			}
		case "3":
			fmt.Println("enter rollno which you want to update")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			fmt.Println("Enter name to update")
			name = getInput()
			updateData(db, rollno, name)
		case "4":
			fmt.Println("enter rollno which you want to delete")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			deleteData(db, rollno)
		case "5":
			os.Exit(0)
		}
	}
}
