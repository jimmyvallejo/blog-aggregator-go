package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/utils"
)

type createUserRequest struct {
	Name string `json:"name"`
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := createUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      request.Name,
	}
	user, err := h.DB.CreateUser(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (h *Handlers) GetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	token, err := utils.ExtractToken(r, "ApiKey ")
	if err != nil {
		if tokenErr, ok := err.(*utils.TokenError); ok {
			respondWithError(w, tokenErr.Code, tokenErr.Message)
		} else {
			respondWithError(w, http.StatusBadRequest, "Invalid API key")
		}
		return
	}

	user, err := h.DB.GetUserByApiKey(r.Context(), token)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			log.Printf("Error getting user by API key: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
