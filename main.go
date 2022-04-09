package main

import (
	"database/sql"
	"errors"
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

	//Insert Person with name
	// err = AddPerson("pek001")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Update Person by id
	// err = UpdatePerson(Person{5, "pek0002233"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Delete Person by id
	// err = DeletePerson(5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Get all person
	persons, err := GetPerson()
	if err != nil {
		panic(err)
	}

	for _, person := range persons {
		fmt.Println(person)
	}

	// person, err := GetPersonWithId(2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print(person)

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

func GetPersonWithId(id int) (*Person, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	sql := db.QueryRow("select person_id, person_name from person where person_id = ?", id)

	person := Person{}

	err = sql.Scan(&person.id, &person.name)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func AddPerson(namePerson string) error {

	query := "insert into person (person_name) values (?)"

	result, err := db.Exec(query, namePerson)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect <= 0 {
		return errors.New("can't insert'")
	}

	return nil
}

func UpdatePerson(person Person) error {

	query := "update person set person_name = ? where person_id = ?"

	result, err := db.Exec(query, person.name, person.id)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect <= 0 {
		return errors.New("can't update'")
	}

	return nil
}

func DeletePerson(id int) error {

	query := "delete from person where person_id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect <= 0 {
		return errors.New("can't delete'")
	}

	return nil
}
