package main

import (
	"fmt"
	"log"
	"net/http"
)

//apiConfig holds the counter for file server hits
type apiConfig struct {
	fileserverHits int
}

func main() {

	// Define the root directory from which files will be served.
	const filepathRoot = "."

	// Define the port on which the server will listen.
	const port = "8080"

	//Initialize the apiConfig struct with fileserverHits counter set to 0.
	apiCfg := apiConfig {
		fileserverHits: 0,
	}

	// Create a new ServeMux, which is a HTTP request multiplexer.
	mux := http.NewServeMux()

	// Register the file server handler with a prefix of "/app".
	// The middlewareMetricsInc middleware will increment the hit counter for each request.
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	// Register the health check endpoint.
	mux.HandleFunc("/healthz", handlerReadiness)

	//Register the metrics endpoint to show the file server hit count.
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)

	// Register the reset endpoint to reset the file server hit count.
	mux.HandleFunc("/reset", apiCfg.handlerReset)

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

// middlewareMetricsInc is a middleware that increments the file server hit counter for each request.
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r) // Call the next handler in the chain.
	})
}

// handleMetrics handles the metrics endpoint and returns how many time the fileserver got hit.
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits)))
}
