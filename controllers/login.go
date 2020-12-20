package controllers

import (
	"encoding/json"
	"net/http"
	"shanyraq/db"
	model "shanyraq/models"
)

var jwtKey = []byte("some secret key")

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			BadRequest(w, r)
			return
		}

		if db.IsValidCredentials(user) {
			//save session return token

		} else {
			BadRequest(w, r)
		}

	} else {
		BadRequest(w, r)
	}
}
