package routes

import (
	"github.com/catalinfl/blog-api/handlers"
	"github.com/go-chi/chi/v5"
)

func CommentsRoute(r *chi.Mux, h *handlers.Handler) {
	r.Get("/comments/{postId}", h.GetCommentsForPost)
	r.Get("/comments/author/{authorId}", h.GetCommentsForAuthor)
	r.Post("/comments", h.CreateComment)
	r.Delete("/comments/{id}", h.DeleteComment)
}
