package main

import (
	"log"
	"net/http"

	"github.com/catalinfl/blog-api/database"
	"github.com/catalinfl/blog-api/handlers"
	"github.com/catalinfl/blog-api/queries"
	"github.com/catalinfl/blog-api/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	conn := database.ConnectDatabase()
	database.CreateTables(conn)
	repo := queries.NewRepo(conn)

	h := handlers.NewHandler(repo)

	routes.AuthorRoute(r, h)
	routes.PostsRoute(r, h)
	routes.CommentsRoute(r, h)

	log.Println("Server started on port :8080")

	http.ListenAndServe(":8080", r)
}
