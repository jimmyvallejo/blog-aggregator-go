package main

import "github.com/jimmyvallejo/blog-aggregator-go/internal/database"

type APIConfig struct {
	Port string
	DB * database.Queries
}
