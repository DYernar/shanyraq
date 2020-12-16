package db

import (
	"database/sql"
	"errors"
	"fmt"
	model "shanyraq/models"

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

//Save user to db
func SaveUser(user model.User) error {
	db, err := DbConnect()
	if err != nil {
		return err
	}

	is_unique, err := IsUniqueUser(user)
	if !is_unique {
		return err
	}

	hashedPassword, err := EncryptPassword([]byte(user.Password))
	if err != nil {
		return err
	}

	_, err = db.Query("insert into users(username, name, surname, email, telephone, password, isValidated) values($1, $2, $3, $4, $5, $6, $7)", user.Username, user.Name, user.Surname, user.Email, user.Telephone, hashedPassword, false)
	db.Close()
	return err
}

//Check if users credentials are already present in a database
func IsUniqueUser(user model.User) (bool, error) {
	db, err := DbConnect()

	if err != nil {
		db.Close()

		return false, err
	}

	err = db.QueryRow(`SELECT username FROM users WHERE username =$1`, user.Username).Scan(&user.Username)

	if err != sql.ErrNoRows {
		db.Close()
		return false, errors.New("username is not unique")
	}

	err = db.QueryRow(`SELECT email FROM users WHERE email =$1`, user.Email).Scan(&user.Email)

	if err != sql.ErrNoRows {
		db.Close()
		return false, errors.New("email is not unique")

	}

	err = db.QueryRow(`SELECT telephone FROM users WHERE telephone =$1`, user.Telephone).Scan(&user.Telephone)

	if err != sql.ErrNoRows {
		db.Close()
		return false, errors.New("telephone is not unique")
	}
	db.Close()
	return true, nil
}

//as its name suggests
func GetUserByUsername(username string) *model.User {
	db, err := DbConnect()

	if err != nil {
		db.Close()
		fmt.Println(err)
		return nil
	}

	row := db.QueryRow(`SELECT userid, username, email, name, surname, telephone, isValidated FROM users WHERE username=$1`, username)
	var user model.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Surname, &user.Telephone, &user.IsValidated)
	if err != nil {
		fmt.Println("Error while getting scanning the user: ", err)
		db.Close()
		return nil
	}
	db.Close()
	return &user
}

//as its name suggests
func GetUserById(id int) *model.User {
	db, err := DbConnect()

	if err != nil {
		db.Close()
		fmt.Println(err)
		return nil
	}

	row := db.QueryRow(`SELECT userid, username, email, name, surname, telephone, isValidated FROM users WHERE userid=$1`, id)
	var user model.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Surname, &user.Telephone, &user.IsValidated)
	if err != nil {
		fmt.Println("Error while getting scanning the user: ", err)
		db.Close()
		return nil
	}
	db.Close()
	return &user
}

//updates user
func UpdateUser(user model.User) error {
	db, err := DbConnect()

	if err != nil {
		return err
	}

	hashedPassword, err := EncryptPassword([]byte(user.Password))

	if err != nil {
		return err
	}

	_, err = db.Exec("update users set username=$1, name=$2, surname=$3, email=$4, telephone=$5, password=$6, isValidated=$7 where userid=$8", user.Username, user.Name, user.Surname, user.Email, user.Telephone, hashedPassword, user.IsValidated, user.ID)

	db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func IsValidCredentials(user model.User) (bool, error) {
	db, err := DbConnect()
	if err != nil {
		db.Close()
		return false, err
	}

	pass := ""

	err = db.QueryRow(`SELECT password FROM users WHERE username = $1`, user.Username).Scan(&pass)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("CHECKING VALID USERNAME ERROR: ", err)
		return false, err
	}

	valid, err := IsValidPassword(user.Password, pass)

	if valid {
		db.Close()
		return true, nil
	}

	err = db.QueryRow(`SELECT password FROM users WHERE email = $1`, user.Email).Scan(&pass)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("CHECKING VALID EMAIL ERROR: ", err)
		return false, err
	}

	valid, err = IsValidPassword(user.Password, pass)

	if valid {
		db.Close()
		return true, nil
	}

	err = db.QueryRow(`SELECT password FROM users WHERE telephone = $1`, user.Telephone).Scan(&pass)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("CHECKING VALID TEL ERROR: ", err)
		return false, err
	}

	valid, err = IsValidPassword(user.Password, pass)

	if valid {
		db.Close()
		return true, nil
	}

	db.Close()
	return false, nil

}

func DeleteUserByUsername(username string) bool {
	db, err := DbConnect()

	if err != nil {
		fmt.Println("Deleting user: ", err)
		db.Close()
		return false
	}

	_, err = db.Exec(`DELETE FROM users WHERE username=$1`, username)

	if err != nil {
		db.Close()
		fmt.Println("DELETING ERROR: ", err)
		return false
	}

	return true
}

func ValidateUser(id int) error {
	db, err := DbConnect()

	if err != nil {
		fmt.Println("VALIDATING USER ERROR: ", err)
		return err
	}

	_, err = db.Exec("update users set isValidated=$1", true)

	if err != nil {
		fmt.Println("VALIDATING USER ERROR: ", err)
		return err
	}

	return nil
}
