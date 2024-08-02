package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
)

type DBInterface interface {
	CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error)
	CreateFeed(ctx context.Context, params database.CreateFeedParams) (database.Feed, error)
	CreateFeedFollow(ctx context.Context, params database.CreateFeedFollowParams) (database.FeedFollow, error)
	GetAllFeeds(ctx context.Context) ([]database.Feed, error)
	RemoveFeedFollow(ctx context.Context, id uuid.UUID) error
	GetAllFeedFollows(ctx context.Context, userID uuid.UUID) ([]database.FeedFollow, error)
	GetUserByApiKey(ctx context.Context, apiKey string) (database.User, error)
}

type Handlers struct {
	DB DBInterface
}

func NewHandlers(db DBInterface) *Handlers {
	return &Handlers{
		DB: db,
	}
}
