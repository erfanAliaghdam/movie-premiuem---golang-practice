package routes

import (
	"movie_premiuem/core/handler"
	"movie_premiuem/core/handler/license_handlers"
	customMiddlewares "movie_premiuem/core/middleware"
	"movie_premiuem/user/handler/auth_handlers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(customMiddlewares.LogRequestsMiddleware) // Add the custom log middleware

	// routes
	mux.Get("/", handler.Index)
	//licenses
	mux.Get("/licenses", license_handlers.LicenseListHandler)
	// auth
	mux.Post("/auth/register", auth_handlers.RegisterUserHandler)

	return mux
}