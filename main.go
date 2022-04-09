package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	// Database Config
	cfg := getDBConfig()
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	persons, err := GetPerson()
	if err != nil {
		panic(err)
	}

	for _, person := range persons {
		fmt.Println(person)
	}

}

func getDBConfig() mysql.Config {
	return mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "learning",
	}
}

func setDatabaseConnection(db *sql.DB) {
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

type Person struct {
	id   int
	name string
}

func GetPerson() ([]Person, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	// Set Datbase Connection
	setDatabaseConnection(db)

	query := "select person_id,person_name from person"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	persons := []Person{}
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.id, &person.name)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, err
}
