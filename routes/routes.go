package routes

import (
	"movie_premiuem/handler"
	"movie_premiuem/handler/license_handlers"
	"movie_premiuem/logger"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(logger.LogRequestsMiddleware) // Add the custom log middleware

	// routes
	mux.Get("/", handler.Index)
	mux.Get("/licenses", license_handlers.LicenseListHandler)

	return mux
}
