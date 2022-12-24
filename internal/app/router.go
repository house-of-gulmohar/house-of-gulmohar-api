package app

import (
	"house-of-gulmohar/internal/e"
	cm "house-of-gulmohar/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) InitRouter() *chi.Mux {
	r := chi.NewRouter()
	// TODO: create custom request logger
	r.Use(middleware.Logger)
	r.Use(cm.HeaderMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api/v1", func(r chi.Router) {
		// products routs
		r.Route("/products", func(r chi.Router) {
			r.Get("/", e.HandleException(s.Product.HandleGetAllProducts))
			r.Get("/{id}", e.HandleException(s.Product.HandleGetProduct))
		})

		// categories routes
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", e.HandleException(s.Category.handleGetAllCategories))
			r.Get("/{id}", e.HandleException(s.Category.handleGetCategory))
		})
	})
	return r
}
