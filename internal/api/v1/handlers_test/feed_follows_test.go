package handlers_test

import (
	"context"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
)

func (m *MockDB) CreateFeedFollow(ctx context.Context, params database.CreateFeedFollowParams) (database.FeedFollow, error) {
	return database.FeedFollow{}, nil
}

func (m *MockDB) GetAllFeedFollows(ctx context.Context, userID uuid.UUID) ([]database.FeedFollow, error) {
	return []database.FeedFollow{}, nil
}

func (m *MockDB) RemoveFeedFollow(ctx context.Context, id uuid.UUID) error {
	return nil
}
