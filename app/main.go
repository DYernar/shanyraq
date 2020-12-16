package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shanyraq/controllers"

	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.MainPage)
	mux.HandleFunc("/auth", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
	fmt.Println("Running on port: " + port)
}
