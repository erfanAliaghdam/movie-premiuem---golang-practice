package routes

import (
	"movie_premiuem/handler"
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

	// routes
	mux.Get("/", handler.Index)

	return mux
}
