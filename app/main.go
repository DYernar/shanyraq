package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	mainControllers "shanyraq/controllers/main"

	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainControllers.MainPage)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
	fmt.Println("Running on port: " + port)
}
