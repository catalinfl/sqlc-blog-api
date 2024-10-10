package routes

import (
	"github.com/catalinfl/blog-api/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthorRoute(r *chi.Mux, h *handlers.Handler) {
	r.Get("/authors/{id}", h.GetAuthor)
	r.Get("/authors", h.GetAllAuthors)
	r.Post("/authors", h.CreateAuthor)
	r.Delete("/authors/{id}", h.DeleteAuthor)

}
