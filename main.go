package main

import (
	"log"
	"net/http"

	"github.com/Paintxd/compassitoMail/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mail", handlers.Create)

	println("Mailservice listening port 8000")
	err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}
