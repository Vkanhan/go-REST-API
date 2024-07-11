package main

import (
	"fmt"
	"net/http"
)

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
