package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/api/common"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
)

type AddFollowRequest struct {
	FeedID uuid.UUID `json:"feed_id"`
}

func (h *Handlers) AddFeedFollow(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(common.UserContextKey).(database.User)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unathorized")
		return
	}

	request := AddFollowRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    request.FeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feedFollow, err := h.DB.CreateFeedFollow(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to add follow")
		return
	}
	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (h *Handlers) RemoveFeedFollow(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/v1/feed_follows/")

	id, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID")
		return
	}

	err = h.DB.RemoveFeedFollow(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found to delete")
		return
	}

	respondNoBody(w, http.StatusOK)
}

func (h *Handlers) GetAllFeedFollows(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(common.UserContextKey).(database.User)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unathorized")
		return
	}

	feedFollows, err := h.DB.GetAllFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Feed follows not found for user")
		return
	}
	respondWithJSON(w, http.StatusOK, feedFollows)
}
