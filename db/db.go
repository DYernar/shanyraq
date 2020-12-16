package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func DbConnect() (*sql.DB, error) {
	configs := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", configs)

	if err != nil {
		fmt.Println("opening database error: ", err)
		return db, err
	}

	_, err = db.Exec("create table if not exists users(userid serial primary key, username varchar, name varchar, surname varchar, email varchar, telephone varchar, password varchar, isValidated bool)")

	if err != nil {
		fmt.Println("table creation error : ", err)
		return db, err
	}

	return db, nil
}

func Drop() {
	db, _ := DbConnect()
	db.Exec("DROP TABLE users")
	db.Close()
}
