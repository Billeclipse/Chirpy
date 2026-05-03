package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Serve files from current directory
	fileServer := http.FileServer(http.Dir("."))

	mux.Handle("/", fileServer)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Serving files on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}