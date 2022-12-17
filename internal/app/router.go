package app

import (
	"house-of-gulmohar/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *Server) InitRouter() *chi.Mux {
	r := chi.NewRouter()
	// TODO: create custom request logger
	// r.Use(middleware.Logger)
	r.Use(middleware.HeaderMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return r
}
