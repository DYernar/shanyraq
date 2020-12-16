package model

type User struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Telephone string `json:"telephone"`
	Email string `json:"email"`
	IsValidated bool `json:"is_validated"`
}
