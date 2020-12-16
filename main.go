package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shanyraq/controllers"
	"shanyraq/db"

	"github.com/rs/cors"
)

func main() {

	_, err := db.DbConnect()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.MainPage)
	mux.HandleFunc("/auth", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)

	handler := cors.Default().Handler(mux)

	fmt.Println("Running on port: " + port)
	err = http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
}
