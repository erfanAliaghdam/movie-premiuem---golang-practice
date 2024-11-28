package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func LogRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()

		// Use a response writer wrapper to capture the status code
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		// Log request details
		log.Printf(
			"Started %s %s - Completed %d %s in %v",
			r.Method,
			r.URL.Path,
			ww.Status(),
			http.StatusText(ww.Status()),
			time.Since(start),
		)

	})
}
