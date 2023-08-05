package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thimc/go-sqlc-demo/sqlc"
)

//go:embed schema.sql
var ddl string

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// Creates the table(s)
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(db)

	// List all the author(s)
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ListAuthors: %+v\n", authors)

	// Create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, sqlc.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: sql.NullString{
			String: "Co-author of The C Programming Language and The Go Programming Language",
			Valid:  true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CreateAuthor: %+v\n", insertedAuthor)

	// Get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compare GetAuthor and the inserted:", reflect.DeepEqual(insertedAuthor, fetchedAuthor))
}
