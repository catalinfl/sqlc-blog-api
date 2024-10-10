package routes

import (
	"github.com/catalinfl/blog-api/handlers"
	"github.com/go-chi/chi/v5"
)

func PostsRoute(r *chi.Mux, h *handlers.Handler) {
	r.Get("/posts/{id}", h.GetPost)
	r.Get("/posts", h.GetAllPosts)
	r.Post("/posts", h.CreatePost)
	r.Put("/posts/{id}", h.EditPost)
	r.Delete("/posts/{id}", h.DeletePost)
}
