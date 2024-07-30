package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jimmyvallejo/blog-aggregator-go/internal/database"
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
