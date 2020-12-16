package controllers

import (
	"fmt"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
