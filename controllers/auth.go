package controllers

import (
	"fmt"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AuthPage")
}
