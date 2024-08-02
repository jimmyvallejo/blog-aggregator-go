package handlers_test

import (
	"context"

	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
)

func (m *MockDB) CreateFeed(ctx context.Context, params database.CreateFeedParams) (database.Feed, error) {
	return database.Feed{}, nil
}

func (m *MockDB) GetAllFeeds(ctx context.Context) ([]database.Feed, error) {
	return []database.Feed{}, nil
}
