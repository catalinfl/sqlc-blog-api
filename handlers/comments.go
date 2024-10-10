package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catalinfl/blog-api/queries"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment queries.CreateCommentParams
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	commentReturned, err := h.Repo.CreateComment(ctx, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentReturned)
}

func (h *Handler) GetCommentsForAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "authorId")
	ctx := r.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	comments, err := h.Repo.GetCommentsForAuthor(ctx, int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "postId")
	ctx := r.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	comments, err := h.Repo.GetCommentsForPost(ctx, int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = h.Repo.DeleteComment(ctx, int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
