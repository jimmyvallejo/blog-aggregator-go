package handlers

import "github.com/jimmyvallejo/blog-aggregator-go/internal/database"

type Handlers struct {
	DB * database.Queries
}

func NewHandlers(db *database.Queries) *Handlers {
	return &Handlers{
		DB: db,
	}
}