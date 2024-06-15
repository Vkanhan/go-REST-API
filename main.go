package main

import (
	"log"
	"net/http"
)

func main() {

	// Define the root directory from which files will be served.
	const filepathRoot = "."

	// Define the port on which the server will listen.
	const port = "8080"

	// Create a new ServeMux, which is a HTTP request multiplexer.
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Create an HTTP server configuration.
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	// Start the HTTP server to listen on the specified port.
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error listening to port 8080")
	}

}


