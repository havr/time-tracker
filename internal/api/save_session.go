package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/havr/time-tracker/internal/models"
)

type saveSessionRequest struct {
	Name     string    `json:"name"`
	Start    time.Time `json:"startTime"`
	Duration int       `json:"duration"`
}

type saveSessionResponse struct {
	ID string `json:"id"`
}

func (h *handler) saveSession(w http.ResponseWriter, r *http.Request) {
	var request saveSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := uuid.New()

	if err := h.sessionStore.SaveSession(&models.Session{
		ID:        uuid.New(),
		Name:      request.Name,
		StartTime: request.Start,
		Duration:  time.Duration(request.Duration) * time.Second,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(saveSessionResponse{
		ID: id.String(),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
