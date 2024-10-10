package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catalinfl/blog-api/queries"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo *queries.Repo
}

func NewHandler(repo *queries.Repo) *Handler {
	return &Handler{
		Repo: repo,
	}
}

func (h *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	author, err := h.Repo.GetAuthor(ctx, int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(author)
}

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author queries.CreateAuthorParams
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	authorReturned, err := h.Repo.CreateAuthor(ctx, author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authorReturned)
}

func (h *Handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authors, err := h.Repo.GetAllAuthors(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authors)

}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteAuthor(ctx, int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Author deleted")
}
