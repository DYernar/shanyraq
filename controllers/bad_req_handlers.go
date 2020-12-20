package controllers

import (
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	err := struct {
		Error string `json:"error"`
	}{
		"page not found",
	}
	vals, _ := json.Marshal(err)
	w.Write(vals)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)

	err := struct {
		Error string `json:"error"`
	}{
		"bad request!",
	}
	vals, _ := json.Marshal(err)
	w.Write(vals)
}

func NotAuthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)

	err := struct {
		Error string `json:"error"`
	}{
		"you are not permitted to see this page",
	}
	vals, _ := json.Marshal(err)
	w.Write(vals)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	err := struct {
		Error string `json:"error"`
	}{
		"internal server",
	}

	vals, _ := json.Marshal(err)
	w.Write(vals)
}
