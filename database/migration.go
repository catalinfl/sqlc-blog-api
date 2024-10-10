package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTables(conn *pgxpool.Pool) {
	authorTable := `CREATE TABLE IF NOT EXISTS authors (
		id BIGSERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	postsTable := `CREATE TABLE IF NOT EXISTS posts (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		author_id BIGINT NOT NULL,
		CONSTRAINT fk_posts_author FOREIGN KEY (author_id) REFERENCES authors(id)
	);`

	commentsTable := `CREATE TABLE IF NOT EXISTS comments (
		id BIGSERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		author_id BIGINT NOT NULL,
		post_id BIGINT NOT NULL,
		CONSTRAINT fk_comments_author FOREIGN KEY (author_id) REFERENCES authors(id),
		CONSTRAINT fk_comments_post FOREIGN KEY (post_id) REFERENCES posts(id)
	);`

	_, err := conn.Exec(context.Background(), authorTable)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), postsTable)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(context.Background(), commentsTable)
	if err != nil {
		panic(err)
	}
}
