package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/api/v1/handlers"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockDB := new(MockDB)
	h := &handlers.Handlers{DB: mockDB}

	reqBody := `{"name": "Test User"}`
	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(reqBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	now := time.Now()

	mockDB.On("CreateUser", mock.Anything, mock.AnythingOfType("database.CreateUserParams")).Return(database.User{
		ID:        uuid.New(),
		Name:      "Test User",
		CreatedAt: now,
		UpdatedAt: now,
	}, nil)

	h.CreateUser(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response database.User
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "Test User", response.Name)
	assert.WithinDuration(t, now, response.CreatedAt, time.Second)
	assert.WithinDuration(t, now, response.UpdatedAt, time.Second)

	mockDB.AssertExpectations(t)
}

func (m *MockDB) GetUserByApiKey(ctx context.Context, apiKey string) (database.User, error) {
	return database.User{}, nil
}
