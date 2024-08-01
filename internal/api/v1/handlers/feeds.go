package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/api/common"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
)

type createFeedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (h *Handlers) CreateFeed(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(common.UserContextKey).(database.User)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unathorized")
		return
	}

	request := createFeedRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      request.Name,
		Url:       request.Url,
		UserID:    user.ID,
	}

	feed, err := h.DB.CreateFeed(r.Context(), feedParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feedFollow, err := h.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed follow")
	}

	response := CreateFeedResponse{
		Feed:       feed,
		FeedFollow: feedFollow,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

func (h *Handlers) GetAllFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := h.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get feeds")
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
