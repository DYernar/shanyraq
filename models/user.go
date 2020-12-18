package model

import "fmt"

type User struct {
	ID          interface{} `bson:"_id"  json:"id,omitempty"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Username    string      `json:"username"`
	Password    string      `json:"password"`
	Telephone   string      `json:"telephone"`
	Email       string      `json:"email"`
	IsValidated bool        `json:"is_validated"`
}

func (u User) PrintUser() {
	res := fmt.Sprintf("ID = %s, Username = %s, Name = %s, Surname = %s, Email = %s, Tel = %s, Password = %s", u.ID, u.Username, u.Name, u.Surname, u.Email, u.Telephone, u.Password)
	fmt.Println(res)
}
