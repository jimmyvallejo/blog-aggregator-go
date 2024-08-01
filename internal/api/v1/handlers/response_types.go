package handlers

import "github.com/jimmyvallejo/blog-aggregator-go/internal/database"

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateFeedResponse struct {
	Feed database.Feed `json:"feed"`
	FeedFollow database.FeedFollow `json:"feed_follow"`
}
