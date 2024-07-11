package main

import (
	"log"
	"net/http"
)

// apiConfig holds the counter for file server hits
type apiConfig struct {
	fileserverHits int
}

func main() {

	// Define the root directory from which files will be served.
	const filepathRoot = "."

	// Define the port on which the server will listen.
	const port = "8080"

	//Initialize the apiConfig struct with fileserverHits counter set to 0.
	apiCfg := apiConfig{
		fileserverHits: 0,
	}

	// Create a new ServeMux, which is a HTTP request multiplexer.
	mux := http.NewServeMux()

	// Register the file server handler with a prefix of "/app".
	// The middlewareMetricsInc middleware will increment the hit counter for each request.
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	// Register the health check endpoint.
	mux.HandleFunc("GET /healthz", handlerReadiness)

	//Register the metrics endpoint to show the file server hit count.
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)

	// Register the reset endpoint to reset the file server hit count.
	mux.HandleFunc("GET /reset", apiCfg.handlerReset)

	// Create an HTTP server configuration.
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux, //ServeMux will handle the incoming requests.
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	// Start the HTTP server to listen on the specified port.
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error listening to port 8080")
	}

}
