package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(startTime)

		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		log.Printf("Response sent in %s\n", duration)
	})
}
