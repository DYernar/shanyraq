package controllers

import (
	"encoding/json"
	"net/http"
	"shanyraq/db"
	model "shanyraq/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("some secret key")
var secretAppToken = "ShanyraqToken"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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

			expirationDate := time.Now().Add(5 * time.Minute)
			claims := &Claims{
				Username: user.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationDate.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				InternalServerError(w, r)
			}

			db.InsertToken(user, tokenString)

			user.Password = ""
			res := struct {
				Status int        `json:"status"`
				Token  string     `json:"token"`
				User   model.User `json:"user"`
			}{
				200,
				tokenString,
				user,
			}

			json, err := json.Marshal(res)

			if err != nil {
				InternalServerError(w, r)
			}

			w.WriteHeader(200)
			w.Write([]byte(string(json)))

		} else {
			BadRequest(w, r)
		}

	} else {
		BadRequest(w, r)
	}
}
