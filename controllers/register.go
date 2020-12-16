package controllers

import (
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else {
		w.WriteHeader(500)
		_, _ = fmt.Fprintf(w, "Bad Request...")
	}
}
